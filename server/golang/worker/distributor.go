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

func (c *DeliveryTaskClient) DeliveryForgetPasswordTask(email string, token string) error {
	send_email_forget_password, err := NewForgetPasswordRequestTaskPayload(
		ForgetPasswordTaskPayload{Email: email, Token: token},
	)
	if err != nil {
		fmt.Printf("Error when adding forget password request task: %v\n", err)
	}
	task_info, err := c.client.Enqueue(send_email_forget_password, asynq.MaxRetry(5))

	if err != nil {
		fmt.Println("Error enqueueing task :", err)
		return err
	}
	log.Info().Str("Forget password Task created", task_info.ID).Str("Queue", task_info.Queue).Msg(":>>>>>>>>>>>>>>>>.")
	return nil
}

func (c *DeliveryTaskClient) DeliveryNewPasswordTask(email string, token string) error {
	send_email_forget_password, err := NewPasswordDeliveryTaskPayload(
		NewPasswordTaskPayload{Email: email},
	)
	if err != nil {
		fmt.Printf("Error when adding new password delivery task: %v\n", err)
	}
	task_info, err := c.client.Enqueue(send_email_forget_password, asynq.MaxRetry(5))

	if err != nil {
		fmt.Println("Error enqueueing task :", err)
		return err
	}
	log.Info().Str("New password delivery Task created", task_info.ID).Str("Queue", task_info.Queue).Msg(":>>>>>>>>>>>>>>>>.")
	return nil
}
