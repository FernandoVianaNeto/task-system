package kafka_pkg

import (
	configs "task-system/cmd/config"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConsumer(topic string) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{configs.KafkaCfg.BrokersHost},
		Topic:   topic,
	})

	return reader
}
