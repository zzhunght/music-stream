package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/pquerna/otp/totp"
	"github.com/rs/zerolog/log"
)

var TypeEmailDeliveryTask = "task:send_email_verify"

type EmailDeliveryTaskPayload struct {
	Email string `json:"email"`
}
type CreateAccountTemp struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	SecretKey string `json:"secret_key"`
}

func NewEmailDeliveryTaskPayload(payload EmailDeliveryTaskPayload) (*asynq.Task, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDeliveryTask, data), nil
}

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
