package api

import (
	"fmt"
	"music-app-backend/internal/app/helper"
	"music-app-backend/sqlc"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pquerna/otp/totp"
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
	err := c.BindJSON(&requestBody)
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
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
	otp, _ := totp.GenerateCode(secret.Secret(), time.Now())
	form := sqlc.CreateAccountParams{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: string(hashedPassword),
		SecretKey: pgtype.Text{
			String: secret.Secret(),
			Valid:  true,
		},
	}
	s.mailsender.SendMailOTP(form.Email, "Mã OTP xác thực của bạn là : "+otp)
	new_acc, err := s.store.CreateAccount(c, form)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, gin.H{
			"error": "Some thing went wrong",
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

	err := c.BindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
	}
	if requestBody.Email == "" || requestBody.Otp == "" {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
	}

	key, err := s.store.GetSecretKey(c, requestBody.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	print("otp ", requestBody.Otp)
	print("key ", key.String)
	valid := totp.Validate(requestBody.Otp, key.String)
	fmt.Print("valid", valid)
	if valid {
		c.JSON(200, gin.H{
			"message": "Xác thực thành công",
			"data":    true,
		})
	}

}

func (s *Server) ResendOTP(c *gin.Context) {
	var requestBody ResendOTPRequest

	err := c.BindJSON(&requestBody)
	if err != nil {
		// Xử lý lỗi nếu có
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
	}
	if requestBody.Email == "" {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
	}

	key, err := s.store.GetSecretKey(c, requestBody.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	otp, _ := totp.GenerateCode(key.String, time.Now())
	valid := totp.Validate(otp, key.String)
	fmt.Println("OTP", otp)
	fmt.Println("KEY", key.String)
	fmt.Println("valid", valid)
	s.mailsender.SendMailOTP(requestBody.Email, "Mã OTP xác thực của bạn là : "+otp)
	if valid {
		c.JSON(200, gin.H{
			"message": "Xác thực thành công",
			"data":    true,
		})
	}

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
	}
	access_token, _ := helper.CreateToken(acc.Email)
	c.JSON(200, gin.H{
		"message": "Login successful",
		"data":    acc,
		"token":   access_token,
	})
}
