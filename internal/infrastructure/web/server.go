package web

import (
	"context"
	domain_usecase "task-system/internal/domain/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router            *gin.Engine
	CreateTaskUsecase domain_usecase.CreateTaskUseCaseInterface
	CreateUserUsecase domain_usecase.CreateUserUsecaseInterface
	GetUserUsecase    domain_usecase.GetUserUsecaseInterface
}

func NewServer(
	ctx context.Context,
	createTaskUsecase domain_usecase.CreateTaskUseCaseInterface,
	createUserUsecase domain_usecase.CreateUserUsecaseInterface,
	getUserUsecase domain_usecase.GetUserUsecaseInterface,
) *Server {
	router := gin.Default()

	server := &Server{
		CreateTaskUsecase: createTaskUsecase,
		CreateUserUsecase: createUserUsecase,
		GetUserUsecase:    getUserUsecase,
	}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
