package web

import (
	"net/http"
	"task-system/internal/domain/dto"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateUserHandler(ctx *gin.Context) {
	var req dto.CreateUserDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Body not compatible with expected"})
		return
	}

	err := s.CreateUserUsecase.Execute(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	// headers := []kafka.Header{
	// 	{Key: "EventType", Value: []byte("task-status-updated")},
	// }

	// err = s.KafkaProducer.WriteMessages(nil, kafka.Message{
	// 	Key:     []byte(task.TaskUuid),
	// 	Headers: headers,
	// 	Value:   messageBytes,
	// })
	// if err != nil {
	// 	return err
	// }

	ctx.Status(http.StatusOK)
}
