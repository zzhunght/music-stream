package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"music-app-backend/internal/app/config"
	"music-app-backend/internal/app/helper"
	"music-app-backend/internal/app/response"
	"music-app-backend/internal/app/services"
	"music-app-backend/pkg/middleware"
	"music-app-backend/sqlc"
	"music-app-backend/worker"
	"net/http"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pquerna/otp/totp"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type VerifyOTPRequest struct {
	Otp   string `json:"otp"`
	Email string `json:"email"`
}
type ResendOTPRequest struct {
	Email string `json:"email"`
}

type LoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type LoginResponse struct {
	SessionID    string       `json:"session_id"`
	User         UserResponse `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
}

type CreateAccountTemp struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	SecretKey string `json:"secret_key"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	Password    string `json:"password"`
}

type UserController struct {
	userService *services.UserService
	rdb         *redis.Client
	tokenMaker  *helper.Token
	taskClient  *worker.DeliveryTaskClient
	config      *config.Config
}

func NewUserController(
	userService *services.UserService,
	rdb *redis.Client,
	tokenMaker *helper.Token,
	taskClient *worker.DeliveryTaskClient,
	config *config.Config,

) *UserController {
	return &UserController{
		userService: userService,
		rdb:         rdb,
		tokenMaker:  tokenMaker,
		taskClient:  taskClient,
		config:      config,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var requestBody RegisterRequest
	// Đọc dữ liệu từ body của request và gán vào biến requestBody
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	if requestBody.Email == "" || requestBody.Password == "" || requestBody.Name == "" {
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	_, err = c.userService.CheckEmailExists(ctx, requestBody.Email)

	if err != pgx.ErrNoRows {
		ctx.JSON(http.StatusConflict, response.ErrorResponse(errors.New("email already exists")))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		ctx.JSON(400, response.ErrorResponse(err))
		return
	}
	// tạo secret key
	secret, errotp := totp.Generate(totp.GenerateOpts{
		Issuer:      "Hung vip pro",
		AccountName: requestBody.Email,
	})
	if errotp != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(errors.New("có lỗi xảy ra trong lúc tạo tài khoản")))
		return
	}
	form := CreateAccountTemp{
		Name:      requestBody.Name,
		Email:     requestBody.Email,
		Password:  string(hashedPassword),
		SecretKey: secret.Secret(),
	}
	key := "register:" + form.Email
	val, err := json.Marshal(form)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(errors.New("có lỗi xảy ra trong lúc tạo tài khoản")))
		return
	}
	err = c.rdb.Set(ctx, key, val, time.Minute*15).Err()

	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(errors.New("có lỗi xảy ra trong lúc tạo tài khoản")))
		return
	}
	err = c.taskClient.DeliveryEmailTaskTask(form.Email)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(errors.New("có lỗi xảy ra trong lúc tạo tài khoản")))
		return
	}
	// otp, _ := totp.GenerateCode(secret.Secret(), time.Now())

	ctx.JSON(201, response.SuccessResponse(true, "Vui lòng kiểm tra email và xác thực OTP để hoàn tất quá trình tạo tài khoản"))

}

func (c *UserController) ChangePassword(ctx *gin.Context) {
	var requestBody ChangePasswordRequest
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(middleware.AuthenticationPayload)
	// authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*helper.TokenPayload)

	// Đọc dữ liệu từ body của request và gán vào biến requestBody
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	acc, err := c.userService.GetAccount(ctx, authPayload.Email)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("not authenticated")))
		return
	}
	validate := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(requestBody.OldPassword))
	if validate != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("incorrect username or password")))
		return
	}
	new_password, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(errors.New("internal error: "+err.Error())))
		return
	}

	err = c.userService.ChangePassword(ctx, sqlc.ChangePasswordParams{
		Email:    authPayload.Email,
		Password: string(new_password),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(errors.New("internal error: "+err.Error())))
		return
	}
	ctx.JSON(200, response.SuccessResponse(true, "Change password successfully"))
}

func (c *UserController) VerifyOTP(ctx *gin.Context) {
	var requestBody VerifyOTPRequest

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(400, response.ErrorResponse(errors.New("invalid request body")))
		return
	}
	var tempUser CreateAccountTemp
	key := "register:" + requestBody.Email
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(400, response.ErrorResponse(errors.New("vui lòng đăng ký trước khi gửi otp")))
		return
	}

	err = json.Unmarshal([]byte(val), &tempUser)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(400, response.ErrorResponse(errors.New("có lỗi xảy ra vui lòng thử lại")))
		return
	}
	fmt.Println("key:", tempUser.SecretKey)
	fmt.Println("OTP:", requestBody.Otp)

	valid := totp.Validate(requestBody.Otp, tempUser.SecretKey)
	log.Info().Any("valid", valid).Msg("")
	if valid {
		newu, err := c.userService.CreateAccount(ctx, sqlc.CreateAccountParams{
			Name:     tempUser.Name,
			Email:    tempUser.Email,
			Password: tempUser.Password,
			SecretKey: pgtype.Text{
				String: tempUser.SecretKey,
				Valid:  true,
			},
		})

		if err != nil {
			ctx.JSON(401, response.ErrorResponse(err))
			return
		}
		ctx.JSON(200, gin.H{
			"message": "Xác thực thành công",
			"data":    newu,
		})
		return
	}
	ctx.JSON(401, response.ErrorResponse(errors.New("xác thực thất bại")))
}

func (c *UserController) ResendOTP(ctx *gin.Context) {
	var requestBody ResendOTPRequest
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(errors.New("invalid request body")))
		return
	}
	c.taskClient.DeliveryEmailTaskTask(requestBody.Email)
	ctx.JSON(200, gin.H{
		"message": "Một OTP đã được gửi đến email của bạn",
		"data":    true,
	})

}

func (c *UserController) ForgetPasswordRequest(ctx *gin.Context) {
	email := ctx.Query("email")
	_, err := mail.ParseAddress(email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(errors.New("email is required")))
		return
	}
	acc, err := c.userService.GetAccount(ctx, email)
	if err != nil {
		// Xử lý lỗi nếu có
		if err != pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse(errors.New("tài khoản không tồn tại")))
			return
		}
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(errors.New("invalid request body")))
		return
	}

	token, _, err := c.tokenMaker.CreateToken(acc.Email, acc.ID, acc.Role, time.Second*60)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(errors.New(err.Error())))
		return
	}
	c.taskClient.DeliveryForgetPasswordTask(acc.Email, token)
	ctx.JSON(http.StatusOK, response.SuccessResponse(true, "Vui lòng kiểm tra Email và làm theo hướng dẫn"))

}
func (c *UserController) ForgetPasswordConfirm(ctx *gin.Context) {
	token := ctx.Query("token")

	auth_payload, err := c.tokenMaker.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New(err.Error())))
		return
	}
	acc, err := c.userService.GetAccount(ctx, auth_payload.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(errors.New(err.Error())))
		return
	}
	c.taskClient.DeliveryNewPasswordTask(acc.Email, token)
	ctx.JSON(http.StatusOK, response.SuccessResponse(true, "Mật khẩu mới đã được gửi đến email của bạn"))

}

func (c *UserController) Login(ctx *gin.Context) {
	var requestBody LoginRequest

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse(errors.New("invalid request body")))
		return
	}
	acc, _ := c.userService.GetAccount(ctx, requestBody.Email)

	validate := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(requestBody.Password))
	if validate != nil {
		ctx.JSON(400, response.ErrorResponse(errors.New("incorrect username or password")))
		return
	}
	access_token, _, err := c.tokenMaker.CreateToken(acc.Email, acc.ID, acc.Role, c.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	refresh_token, rf_payload, err := c.tokenMaker.CreateToken(acc.Email, acc.ID, acc.Role, c.config.RefreshTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	session, err := c.userService.CreateSession(ctx, sqlc.CreateSessionParams{
		ID:           rf_payload.ID,
		Email:        rf_payload.Email,
		RefreshToken: refresh_token,
		ExpiredAt: pgtype.Timestamp{
			Time:  rf_payload.ExpiredAt,
			Valid: true,
		},
		ClientAgent: ctx.Request.UserAgent(),
		ClientIp:    ctx.ClientIP(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &LoginResponse{
		SessionID: session.ID.String(),
		User: UserResponse{
			Email:     acc.Email,
			Name:      acc.Name,
			ID:        acc.ID,
			CreatedAt: acc.CreatedAt.Time,
			UpdatedAt: acc.UpdatedAt.Time,
		},
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}

	ctx.JSON(200, gin.H{
		"message": "Login successful",
		"data":    resp,
	})
}

func (c *UserController) GetUser(ctx *gin.Context) {

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey)
	fmt.Print(authPayload)

	ctx.JSON(http.StatusOK, gin.H{
		"data": authPayload,
	})
}

type RenewTokenBody struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (c *UserController) RenewToken(ctx *gin.Context) {

	var body RenewTokenBody
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("token not found")))
		return
	}

	token_payload, err := c.tokenMaker.VerifyToken(body.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(err))
		return
	}
	session, err := c.userService.GetSession(ctx, token_payload.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("Unauthorized")))
		return
	}
	if session.ID != token_payload.ID {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("Unauthorized")))
		return
	}
	if session.Email != token_payload.Email {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("Unauthorized")))
		return
	}

	if session.RefreshToken != body.RefreshToken {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("Unauthorized")))
		return
	}

	acc, err := c.userService.GetAccount(ctx, session.Email)

	if session.RefreshToken != body.RefreshToken {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("Unauthorized")))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("Unauthorized")))
		return
	}

	new_access_token, _, err := c.tokenMaker.CreateToken(acc.Email, acc.ID, acc.Role, c.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse(errors.New("Unauthorized")))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(new_access_token, "Access token created"))
}
