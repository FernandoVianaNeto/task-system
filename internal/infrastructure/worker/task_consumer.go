package workers

import (
	"context"
	"log"
	configs "task-system/cmd/config"
	"task-system/pkg/kafka"
)

func StartTaskStatusUpdatedConsumer(ctx context.Context) {
	consumer := kafka.NewKafkaConsumer(configs.KafkaCfg.TaskStatusUpdatedTopic)

	msg, err := consumer.ReadMessage(ctx)

	if err != nil {
		log.Fatal(err)
	}

	TaskUpdatedStatusHandler(msg)
}
