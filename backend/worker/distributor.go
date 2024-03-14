package worker

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/hibiken/asynq"
)

type DeliveryTaskClient struct {
	client *asynq.Client
}

func NewDeliveryTaskClient(opts asynq.RedisClientOpt) *DeliveryTaskClient {
	return &DeliveryTaskClient{
		client: asynq.NewClient(opts),
	}
}

func (c *DeliveryTaskClient) DeliveryEmailTaskTask(email string) error {
	send_email_verify_task, err := NewEmailDeliveryTaskPayload(
		EmailDeliveryTaskPayload{Email: email},
	)
	if err != nil {
		fmt.Printf("Error when adding email task: %v\n", err)
	}
	task_info, err := c.client.Enqueue(send_email_verify_task, asynq.MaxRetry(5))

	if err != nil {
		fmt.Println("Error enqueueing task :", err)
		return err
	}
	log.Info().Str("Task created", task_info.ID).Str("Queue", task_info.Queue).Msg(":>>>>>>>>>>>>>>>>.")
	return nil
}
