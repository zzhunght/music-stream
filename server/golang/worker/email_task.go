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
	key, err := process.store.GetSecretKey(ctx, p.Email)
	if err != nil {
		log.Err(err).Msg("Failed to get secret key")
	}
	otp, err := totp.GenerateCode(key.String, time.Now().Add(60))

	if err != nil {
		log.Err(err).Msg("Failed to create otp")
	}
	process.mailer.SendMailOTP(p.Email, "Mã OTP xác thực của bạn là : "+otp)

	log.Info().Str("Email : ", p.Email).Msg("Sending Email : >>>>>>>>>>>>>>")

	// Email delivery code ...
	return nil
}
