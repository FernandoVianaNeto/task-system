package workers

import (
	"context"
	"fmt"
	configs "task-system/cmd/config"
	kafka_pkg "task-system/pkg/kafka"
)

func StartTaskStatusUpdatedConsumer(ctx context.Context) {
	fmt.Println(configs.KafkaCfg)

	consumer := kafka_pkg.NewKafkaConsumer(configs.KafkaCfg.TaskStatusUpdatedTopic)

	msg, err := consumer.ReadMessage(ctx)

	fmt.Println(msg, err)

	// for {
	// 	msg, err := consumer.ReadMessage(ctx)

	// 	fmt.Println("MSG", msg)

	// 	if err != nil {
	// 		log.Println("Error reading message:", err)
	// 		continue
	// 	}

	// 	TaskUpdatedStatusHandler(msg)
	// }
}
