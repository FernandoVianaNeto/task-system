package workers

import (
	"context"
	"log"
	configs "task-system/cmd/config"
	kafka_pkg "task-system/pkg/kafka"
)

func StartTaskStatusUpdatedConsumer(ctx context.Context) {
	consumer := kafka_pkg.NewKafkaConsumer(configs.KafkaCfg.TaskStatusUpdatedTopic)

	msg, err := consumer.ReadMessage(ctx)

	if err != nil {
		log.Fatal(err)
	}

	TaskUpdatedStatusHandler(msg)
}
