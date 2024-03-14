package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
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
	log.Printf("Sending Email to User: email=%d", p.Email)
	// Email delivery code ...
	return nil
}
