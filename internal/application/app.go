package app

import (
	"context"
	"fmt"
	configs "task-system/cmd/config"
	usecase "task-system/internal/application/usecases"
	domain_repository "task-system/internal/domain/repository"
	domain_usecase "task-system/internal/domain/usecase"
	"task-system/internal/infrastructure/repository"
	"task-system/internal/infrastructure/web"
	mysql_pkg "task-system/pkg/mysql"

	"gorm.io/gorm"
)

type Usecases struct {
	CreateTaskUsecase domain_usecase.CreateTaskUseCaseInterface
}

type Repositories struct {
	TaskRepository domain_repository.TaskRepositoryInterface
}

func NewApplication() *web.Server {
	ctx := context.Background()

	mysqlConnectionInput := mysql_pkg.MySqlInput{
		User:     configs.MySqlCfg.User,
		Password: configs.MySqlCfg.Password,
		Host:     configs.MySqlCfg.Host,
		Port:     configs.MySqlCfg.Port,
		Name:     configs.MySqlCfg.Name,
	}

	db, err := mysql_pkg.ConnectToDatabase(mysqlConnectionInput)

	if err != nil {
		fmt.Println(err, "Failed to connect to database")
	}

	repositories := NewRepositories(ctx, db)

	usecases := NewUsecases(ctx, repositories)

	srv := web.NewServer(ctx, usecases.CreateTaskUsecase)

	return srv
}

func NewUsecases(ctx context.Context, repositories Repositories) Usecases {
	createTaskUsecase := usecase.NewCreateTaskUsecase(repositories.TaskRepository)

	return Usecases{
		CreateTaskUsecase: createTaskUsecase,
	}
}

func NewRepositories(ctx context.Context, db *gorm.DB) Repositories {
	taskRepository := repository.NewTaskRepository(db)

	return Repositories{
		TaskRepository: taskRepository,
	}
}
