package kafka_pkg

import (
	"fmt"
	configs "task-system/cmd/config"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConsumer(topic string) *kafka.Reader {
	fmt.Println(configs.KafkaCfg.BrokersHost)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{configs.KafkaCfg.BrokersHost},
		Topic:       topic,
		StartOffset: -1,
	})

	return reader
}
