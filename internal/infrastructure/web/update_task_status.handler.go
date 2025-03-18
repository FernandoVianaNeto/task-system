package web

import (
	"encoding/json"
	"net/http"
	"strings"
	configs "task-system/cmd/config"
	"task-system/internal/domain/dto"
	"task-system/internal/infrastructure/web/middleware"
	infra_request "task-system/internal/infrastructure/web/request"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/segmentio/kafka-go"
)

type TaskNotification struct {
	UserUuid  string `json:"user_uuid"`
	NewStatus string `json:"new_status"`
	TaskUuid  string `json:"task_uuid"`
	TaskTitle string `json:"task_title"`
	Timestamp string `json:"timestamp"`
}

func (s *Server) UpdateTaskStatusHandler(ctx *gin.Context) {
	var req infra_request.UpdateTaskStatusRequest
	var updateTaskStatusDto dto.UpdateTaskStatusDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Body not compatible with expected"})
		return
	}

	token := ctx.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)
	parsedToken, err := jwt.ParseWithClaims(token, &middleware.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.ApplicationCfg.PasswordSecretHash), nil
	})

	if claims, ok := parsedToken.Claims.(*middleware.JwtClaims); ok {
		updateTaskStatusDto.UserUuid = claims.Uuid
	}

	updateTaskStatusDto.TaskStatus = req.NewStatus
	updateTaskStatusDto.TaskUuid = req.TaskUuid

	err = s.UpdateTaskStatusUsecase.Execute(ctx, updateTaskStatusDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	s.produceKafkaMessage(ctx, updateTaskStatusDto)

	ctx.Status(http.StatusOK)
}

func (s *Server) produceKafkaMessage(ctx *gin.Context, req dto.UpdateTaskStatusDto) {
	headers := []kafka.Header{
		{Key: "EventType", Value: []byte("task-status-updated")},
	}

	task := TaskNotification{
		UserUuid:  req.UserUuid,
		NewStatus: req.TaskStatus,
		TaskUuid:  req.TaskUuid,
		Timestamp: time.Now().String(),
	}

	messageBytes, _ := json.Marshal(task)

	s.KafkaProducer.WriteMessages(ctx, kafka.Message{
		Key:     []byte(req.TaskUuid),
		Headers: headers,
		Value:   messageBytes,
	})
}
