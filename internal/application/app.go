package app

import (
	"context"
	"fmt"
	"log"
	configs "task-system/cmd/config"
	usecase "task-system/internal/application/usecases"
	domain_repository "task-system/internal/domain/repository"
	domain_usecase "task-system/internal/domain/usecase"
	repository_task "task-system/internal/infrastructure/repository/task"
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

	err = NewMigrations(db)

	if err != nil {
		log.Fatal(err)
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
	taskRepository := repository_task.NewTaskRepository(db)

	return Repositories{
		TaskRepository: taskRepository,
	}
}

func NewTaskMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&repository_task.Task{})
	if err != nil {
		log.Fatal("Could not run migrations':", err)
	}

	return err
}

func NewUserMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&repository_task.Task{})
	if err != nil {
		log.Fatal("Could not run migrations':", err)
	}

	return err
}

func NewMigrations(db *gorm.DB) error {
	err := NewUserMigration(db)

	if err != nil {
		log.Fatal("Could not run migrations':", err)
	}

	err = NewTaskMigration(db)

	if err != nil {
		log.Fatal("Could not run migrations':", err)
	}

	return err
}
