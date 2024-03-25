package api

import (
	"errors"
	"fmt"
	api "music-app-backend/internal/app/api/middleware"
	"music-app-backend/sqlc"
	"net/http"

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

type LoginResponse struct {
	SessionID    any                `json:"session_id"`
	User         sqlc.GetAccountRow `json:"user"`
	AccessToken  string             `json:"access_token"`
	RefreshToken string             `json:"refresh_token"`
}

// ListUsers là handler cho việc liệt kê các users
func ListUsers(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "List of users",
		"data":    "",
	})
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
		c.JSON(409, ErrorResponse(errors.New("email already exists")))
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
		c.JSON(400, gin.H{
			"error": "Some thing went wrong",
		})
		return
	}
	// otp, _ := totp.GenerateCode(secret.Secret(), time.Now())
	form := sqlc.CreateAccountParams{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: string(hashedPassword),
		SecretKey: pgtype.Text{
			String: secret.Secret(),
			Valid:  true,
		},
	}
	// s.mailsender.SendMailOTP(form.Email, "Mã OTP xác thực của bạn là : "+otp)
	arg := sqlc.CreateAccountWithTxParams{
		Params: form,
		AfterFunction: func(email string) error {
			return s.task_client.DeliveryEmailTaskTask(email)
		},
	}
	new_acc, err := s.store.CreateAccountWithTx(c, arg)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, gin.H{
			"error": "Có lỗi xảy ra trong lúc tạo tài khoản",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Tạo tài khoản thành công",
		"data":    new_acc,
	})

}

func (s *Server) VerifyOTP(c *gin.Context) {
	var requestBody VerifyOTPRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, ErrorResponse(errors.New("invalid request body")))
		return
	}
	key, err := s.store.GetSecretKey(c, requestBody.Email)
	log.Info().Str("otp ", requestBody.Otp).Msg("")
	log.Info().Str("key ", key.String).Msg("")
	if err != nil {
		c.JSON(400, ErrorResponse(err))
		return
	}

	valid := totp.Validate(requestBody.Otp, key.String)
	log.Info().Any("valid", valid).Msg("")
	if valid {
		s.store.VerifyAccount(c, requestBody.Email)
		c.JSON(200, gin.H{
			"message": "Xác thực thành công",
			"data":    true,
		})
		return
	}
	c.JSON(401, ErrorResponse(errors.New("Xác thực thất bại")))
}

func (s *Server) ResendOTP(c *gin.Context) {
	var requestBody ResendOTPRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
	}
	s.task_client.DeliveryEmailTaskTask(requestBody.Email)
	c.JSON(200, gin.H{
		"message": "Một OTP đã được gửi đến email của bạn",
		"data":    true,
	})

}

func (s *Server) Login(c *gin.Context) {
	var requestBody LoginRequest

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid login request",
		})
	}

	if requestBody.Email == "" || requestBody.Password == "" {
		c.JSON(400, gin.H{
			"error": "Invalid login request",
		})
	}

	acc, _ := s.store.GetAccount(c, requestBody.Email)

	validate := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(requestBody.Password))
	if validate != nil {
		c.JSON(400, gin.H{
			"error": "Incorrect username or password",
		})
		return
	}
	access_token, _, err := s.token_maker.CreateToken(acc.Email, acc.Role, s.config.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	refresh_token, rf_payload, err := s.token_maker.CreateToken(acc.Email, acc.Role, s.config.RefreshTokenDuration)
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
		SessionID:    session.ID,
		User:         acc,
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
