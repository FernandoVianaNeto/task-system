package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type TaskNotification struct {
	UserUuid  string `json:"user_uuid"`
	NewStatus string `json:"new_status"`
	TaskUuid  string `json:"task_uuid"`
	TaskTitle string `json:task_title`
	Timestamp string `json:"timestamp"`
}

func TaskUpdatedStatusHandler(msg kafka.Message) {
	for {
		var task TaskNotification
		if err := json.Unmarshal(msg.Value, &task); err != nil {
			log.Println("Error decoding task update:", err)
			continue
		}

		headersMap := make(map[string]string)
		for _, header := range msg.Headers {
			headersMap[header.Key] = string(header.Value)
		}

		fmt.Printf("📢 The tech performed the task %s (ID: %s) on %s. Status: %s\n",
			task.TaskTitle, task.TaskUuid, task.Timestamp, task.NewStatus)
	}
}
