package web

import (
	"context"
	domain_usecase "task-system/internal/domain/usecase"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type Server struct {
	router        *gin.Engine
	Usecases      domain_usecase.Usecases
	KafkaProducer *kafka.Writer
}

func NewServer(
	ctx context.Context,
	usecases domain_usecase.Usecases,
	kafkaProducer *kafka.Writer,
) *Server {
	router := gin.Default()

	server := &Server{
		Usecases: domain_usecase.Usecases{
			CreateTaskUsecase:       usecases.CreateTaskUsecase,
			CreateUserUsecase:       usecases.CreateUserUsecase,
			AuthUsecase:             usecases.AuthUsecase,
			ListTaskUsecase:         usecases.ListTaskUsecase,
			UpdateTaskStatusUsecase: usecases.UpdateTaskStatusUsecase,
			DeleteTaskUsecase:       usecases.DeleteTaskUsecase,
		},
		KafkaProducer: kafkaProducer,
	}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
