package api

import (
	"encoding/json"
	"errors"
	"fmt"
	api "music-app-backend/internal/app/api/middleware"
	"music-app-backend/internal/app/helper"
	"music-app-backend/sqlc"
	"net/http"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pquerna/otp/totp"
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

func (s *Server) Register(c *gin.Context) {
	var requestBody RegisterRequest
	// Đọc dữ liệu từ body của request và gán vào biến requestBody
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	if requestBody.Email == "" || requestBody.Password == "" || requestBody.Name == "" {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	_, err = s.store.CheckEmailExists(c, requestBody.Email)

	if err != pgx.ErrNoRows {
		c.JSON(http.StatusConflict, ErrorResponse(errors.New("email already exists")))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		c.JSON(400, ErrorResponse(err))
		return
	}
	// tạo secret key
	secret, errotp := totp.Generate(totp.GenerateOpts{
		Issuer:      "Hung vip pro",
		AccountName: requestBody.Email,
	})
	if errotp != nil {
		// Xử lý lỗi nếu có
		c.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("có lỗi xảy ra trong lúc tạo tài khoản")))
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
		c.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("có lỗi xảy ra trong lúc tạo tài khoản")))
		return
	}
	err = s.rdb.Set(c, key, val, time.Minute*15).Err()

	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("có lỗi xảy ra trong lúc tạo tài khoản")))
		return
	}
	err = s.task_client.DeliveryEmailTaskTask(form.Email)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("có lỗi xảy ra trong lúc tạo tài khoản")))
		return
	}
	// otp, _ := totp.GenerateCode(secret.Secret(), time.Now())

	c.JSON(201, SuccessResponse(true, "Vui lòng kiểm tra email và xác thực OTP để hoàn tất quá trình tạo tài khoản"))

}

func (s *Server) ChangePassword(ctx *gin.Context) {
	var requestBody ChangePasswordRequest
	authPayload := ctx.MustGet(api.AuthorizationPayloadKey).(*helper.TokenPayload)

	// Đọc dữ liệu từ body của request và gán vào biến requestBody
	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	acc, err := s.store.GetAccount(ctx, authPayload.Email)
	if err != nil {
		// Xử lý lỗi nếu có
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("not authenticated")))
		return
	}
	validate := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(requestBody.OldPassword))
	if validate != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New("incorrect username or password")))
		return
	}
	new_password, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("internal error: "+err.Error())))
		return
	}

	err = s.store.ChangePassword(ctx, sqlc.ChangePasswordParams{
		Email:    authPayload.Email,
		Password: string(new_password),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New("internal error: "+err.Error())))
		return
	}
	ctx.JSON(200, SuccessResponse(true, "Change password successfully"))
}

func (s *Server) VerifyOTP(c *gin.Context) {
	var requestBody VerifyOTPRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, ErrorResponse(errors.New("invalid request body")))
		return
	}
	var tempUser CreateAccountTemp
	key := "register:" + requestBody.Email
	val, err := s.rdb.Get(c, key).Result()
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, ErrorResponse(errors.New("vui lòng đăng ký trước khi gửi otp")))
		return
	}

	err = json.Unmarshal([]byte(val), &tempUser)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, ErrorResponse(errors.New("có lỗi xảy ra vui lòng thử lại")))
		return
	}
	fmt.Println("key:", tempUser.SecretKey)
	fmt.Println("OTP:", requestBody.Otp)

	valid := totp.Validate(requestBody.Otp, tempUser.SecretKey)
	log.Info().Any("valid", valid).Msg("")
	if valid {
		newu, err := s.store.CreateAccount(c, sqlc.CreateAccountParams{
			Name:     tempUser.Name,
			Email:    tempUser.Email,
			Password: tempUser.Password,
			SecretKey: pgtype.Text{
				String: tempUser.SecretKey,
				Valid:  true,
			},
		})

		if err != nil {
			c.JSON(401, ErrorResponse(err))
			return
		}
		c.JSON(200, gin.H{
			"message": "Xác thực thành công",
			"data":    newu,
		})
		return
	}
	c.JSON(401, ErrorResponse(errors.New("xác thực thất bại")))
}

func (s *Server) ResendOTP(c *gin.Context) {
	var requestBody ResendOTPRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(http.StatusBadRequest, ErrorResponse(errors.New("invalid request body")))
		return
	}
	s.task_client.DeliveryEmailTaskTask(requestBody.Email)
	c.JSON(200, gin.H{
		"message": "Một OTP đã được gửi đến email của bạn",
		"data":    true,
	})

}

func (s *Server) ForgetPasswordRequest(ctx *gin.Context) {
	email := ctx.Query("email")
	_, err := mail.ParseAddress(email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(errors.New("email is required")))
		return
	}
	acc, err := s.store.GetAccount(ctx, email)
	if err != nil {
		// Xử lý lỗi nếu có
		if err != pgx.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(errors.New("tài khoản không tồn tại")))
			return
		}
		ctx.JSON(http.StatusBadRequest, ErrorResponse(errors.New("invalid request body")))
		return
	}

	token, _, err := s.token_maker.CreateToken(acc.Email, acc.ID, acc.Role, time.Second*60)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New(err.Error())))
		return
	}
	s.task_client.DeliveryForgetPasswordTask(acc.Email, token)
	ctx.JSON(http.StatusOK, SuccessResponse(true, "Vui lòng kiểm tra Email và làm theo hướng dẫn"))

}
func (s *Server) ForgetPasswordConfirm(ctx *gin.Context) {
	token := ctx.Query("token")

	auth_payload, err := s.token_maker.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, ErrorResponse(errors.New(err.Error())))
		return
	}
	acc, err := s.store.GetAccount(ctx, auth_payload.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(errors.New(err.Error())))
		return
	}
	s.task_client.DeliveryNewPasswordTask(acc.Email, token)
	ctx.JSON(http.StatusOK, SuccessResponse(true, "Mật khẩu mới đã được gửi đến email của bạn"))

}

func (s *Server) Login(c *gin.Context) {
	var requestBody LoginRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(400, ErrorResponse(errors.New("invalid request body")))
		return
	}
	acc, _ := s.store.GetAccount(c, requestBody.Email)

	validate := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(requestBody.Password))
	if validate != nil {
		c.JSON(400, ErrorResponse(errors.New("incorrect username or password")))
		return
	}
	access_token, _, err := s.token_maker.CreateToken(acc.Email, acc.ID, acc.Role, s.config.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	refresh_token, rf_payload, err := s.token_maker.CreateToken(acc.Email, acc.ID, acc.Role, s.config.RefreshTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	session, err := s.store.CreateSession(c, sqlc.CreateSessionParams{
		ID:           rf_payload.ID,
		Email:        rf_payload.Email,
		RefreshToken: refresh_token,
		ExpiredAt: pgtype.Timestamp{
			Time:  rf_payload.ExpiredAt,
			Valid: true,
		},
		ClientAgent: c.Request.UserAgent(),
		ClientIp:    c.ClientIP(),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
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

	c.JSON(200, gin.H{
		"message": "Login successful",
		"data":    resp,
	})
}

func (s *Server) GetUser(c *gin.Context) {

	authPayload := c.MustGet(api.AuthorizationPayloadKey)
	fmt.Print(authPayload)

	c.JSON(http.StatusOK, gin.H{
		"data": authPayload,
	})
}
