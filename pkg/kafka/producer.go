package kafka_pkg

import (
	configs "task-system/cmd/config"

	"github.com/segmentio/kafka-go"
)

func NewProducer(topic string) *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{configs.KafkaCfg.BrokersHost},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return writer
}
