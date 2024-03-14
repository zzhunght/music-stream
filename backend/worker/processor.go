package worker

import "github.com/hibiken/asynq"

type ProcessorRedisTasks struct {
	client *asynq.Server
}

func NewProcessorTaskClient(opts asynq.RedisClientOpt) *ProcessorRedisTasks {
	client := asynq.NewServer(opts, asynq.Config{})
	return &ProcessorRedisTasks{
		client: client,
	}
}

func (process *ProcessorRedisTasks) Start() error {

	mux := asynq.NewServeMux()

	mux.HandleFunc(TypeEmailDeliveryTask, process.HandleEmailDeliveryTask)

	return process.client.Start(mux)
}
