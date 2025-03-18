package workers

import (
	"context"
	"log"
	configs "task-system/cmd/config"
	workers "task-system/internal/infrastructure/worker"
	"task-system/pkg/kafka"
)

type TaskNotification struct {
	UserUuid  string `json:"user_uuid"`
	NewStatus string `json:"new_status"`
	TaskUuid  string `json:"task_uuid"`
	TaskTitle string `json:"task_title"`
	Timestamp string `json:"timestamp"`
}

func StartTaskStatusUpdatedProducer(ctx context.Context) {
	consumer := kafka.NewKafkaConsumer(configs.KafkaCfg.TaskStatusUpdatedTopic)

	msg, err := consumer.ReadMessage(ctx)

	if err != nil {
		log.Fatal(err)
	}

	workers.TaskUpdatedStatusHandler(msg)
}
