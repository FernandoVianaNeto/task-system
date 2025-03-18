package workers

import (
	"context"
	"fmt"
	"log"
	configs "task-system/cmd/config"
	kafka_pkg "task-system/pkg/kafka"
)

func StartTaskStatusUpdatedConsumer(ctx context.Context) {
	consumer := kafka_pkg.NewKafkaConsumer(configs.KafkaCfg.TaskStatusUpdatedTopic)

	fmt.Println("Starting Kafka Consumer...")

	for {
		msg, err := consumer.ReadMessage(ctx)

		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		TaskUpdatedStatusHandler(msg)
	}
}
