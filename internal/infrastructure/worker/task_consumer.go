package workers

import (
	"context"
	"log"
	configs "task-system/cmd/config"
	kafka_pkg "task-system/pkg/kafka"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type KafkaConsumerInterface interface {
	ReadMessage(ctx context.Context) ([]byte, error)
}

func StartTaskStatusUpdatedConsumer(ctx context.Context) {
	consumer := kafka_pkg.NewKafkaConsumer(configs.KafkaCfg.TaskStatusUpdatedTopic)

	for {
		msg, err := consumer.ReadMessage(ctx)

		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		TaskUpdatedStatusHandler(msg)
	}
}
