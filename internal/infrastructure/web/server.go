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
	AuthUsecase       domain_usecase.AuthUsecaseInterface
	ListTaskUsecase   domain_usecase.ListTaskUsecaseInterface
}

func NewServer(
	ctx context.Context,
	createTaskUsecase domain_usecase.CreateTaskUseCaseInterface,
	createUserUsecase domain_usecase.CreateUserUsecaseInterface,
	getUserUsecase domain_usecase.GetUserUsecaseInterface,
	authUsecase domain_usecase.AuthUsecaseInterface,
	listTaskUsecase domain_usecase.ListTaskUsecaseInterface,
) *Server {
	router := gin.Default()

	server := &Server{
		CreateTaskUsecase: createTaskUsecase,
		CreateUserUsecase: createUserUsecase,
		GetUserUsecase:    getUserUsecase,
		AuthUsecase:       authUsecase,
		ListTaskUsecase:   listTaskUsecase,
	}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
