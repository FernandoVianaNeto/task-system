package kafka

import (
	"encoding/json"
	"fmt"
	configs "task-system/cmd/config"

	"github.com/segmentio/kafka-go"
)

type TaskNotification struct {
	UserUuid  string `json:"user_uuid"`
	NewStatus string `json:"new_status"`
	TaskUuid  string `json:"task_uuid"`
	Timestamp string `json:"timestamp"`
}

func SendTaskStatusUpdateNotification(task TaskNotification) error {
	topic := configs.KafkaCfg.TaskStatusUpdatedTopic

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{configs.KafkaCfg.BrokersHost},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	messageBytes, err := json.Marshal(task)
	if err != nil {
		return err
	}

	headers := []kafka.Header{
		{Key: "EventType", Value: []byte("task-status-updated")},
	}

	err = writer.WriteMessages(nil, kafka.Message{
		Key:     []byte(task.TaskUuid),
		Headers: headers,
		Value:   messageBytes,
	})
	if err != nil {
		return err
	}

	fmt.Println("task status updated events sent", string(messageBytes))
	return nil
}
