package worker

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	db "music-app-backend/sqlc"
	"time"

	"github.com/hibiken/asynq"
	"github.com/pquerna/otp/totp"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

var TypeEmailDeliveryTask = "task:send_email_verify"
var TypeEmailDeliveryForgetPasswordRequest = "task:send_email_verify_forgetpassword"
var TypeEmailDeliveryNewPassword = "task:send_email_new_password_reset"

type EmailDeliveryTaskPayload struct {
	Email string `json:"email"`
}

type ForgetPasswordTaskPayload struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type NewPasswordTaskPayload struct {
	Email string `json:"email"`
}
type CreateAccountTemp struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	SecretKey string `json:"secret_key"`
}

// new Task interface
func NewEmailDeliveryTaskPayload(payload EmailDeliveryTaskPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDeliveryTask, data), nil
}
func NewPasswordDeliveryTaskPayload(payload NewPasswordTaskPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDeliveryNewPassword, data), nil
}

func NewForgetPasswordRequestTaskPayload(payload ForgetPasswordTaskPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDeliveryForgetPasswordRequest, data), nil
}

func generatePassword(length int) string {
	var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	for i := range bytes {
		max := big.NewInt(int64(len(charset)))
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return ""
		}
		bytes[i] = charset[n.Int64()]
	}
	return string(bytes)
}

// processor Handle

func (process *ProcessorRedisTasks) HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	var tempUser CreateAccountTemp
	key := "register:" + p.Email
	val, err := process.rdb.Get(ctx, key).Result()
	if err != nil {
		log.Err(err).Msg("Failed to get secret key")
	}
	err = json.Unmarshal([]byte(val), &tempUser)
	if err != nil {
		log.Err(err).Msg("Failed to decode json")
	}

	otp, err := totp.GenerateCode(tempUser.SecretKey, time.Now().Add(60))

	if err != nil {
		log.Err(err).Msg("Failed to create otp")
	}
	process.mailer.SendMailOTP(p.Email, "Mã OTP xác thực của bạn là : "+otp)

	log.Info().Str("Email : ", p.Email).Msg("Sending Email : >>>>>>>>>>>>>>")

	// Email delivery code ...
	return nil
}

func (process *ProcessorRedisTasks) HandleEmailDeliveryForgetPasswordRequestTask(ctx context.Context, t *asynq.Task) error {
	var p ForgetPasswordTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	url := "http://localhost:8080/api/v1/user/confirm-forget-password?token=" + p.Token
	process.mailer.SendMail(
		p.Email,
		"Click vào đường link phía dưới để đặt lại mật khẩu mới của bạn (thời gian hết hạn 90 giây) : "+url,
		"Đặt lại mật khẩu",
	)

	log.Info().Str("Email : ", p.Email).Msg("Sending Email Forget Password Request : >>>>>>>>>>>>>>")

	// Email delivery code ...
	return nil
}

func (process *ProcessorRedisTasks) HandleEmailDeliveryNewtPasswordTask(ctx context.Context, t *asynq.Task) error {
	var p NewPasswordTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	new_password := generatePassword(8)

	hash_password, err := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	if err != nil {
		log.Err(err).Msg("Failed to hash password" + err.Error())
	}
	process.store.ChangePassword(ctx, db.ChangePasswordParams{
		Password: string(hash_password),
		Email:    p.Email,
	})
	process.mailer.SendMail(
		p.Email,
		"Mật khẩu mới của bạn là : "+new_password,
		"Cấp lại mật khẩu mới",
	)

	log.Info().Str("Send New Password to Email : ", p.Email).Msg("Sending New Pasword to Email : >>>>>>>>>>>>>>")

	// Email delivery code ...
	return nil
}
