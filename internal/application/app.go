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
	repository_user "task-system/internal/infrastructure/repository/user"
	"task-system/internal/infrastructure/web"
	kafka_pkg "task-system/pkg/kafka"
	mysql_pkg "task-system/pkg/mysql"

	"gorm.io/gorm"
)

type Usecases struct {
	CreateTaskUsecase       domain_usecase.CreateTaskUseCaseInterface
	CreateUserUsecase       domain_usecase.CreateUserUsecaseInterface
	GetUserUsecase          domain_usecase.GetUserUsecaseInterface
	AuthUsecase             domain_usecase.AuthUsecaseInterface
	ListTaskUsecase         domain_usecase.ListTaskUsecaseInterface
	UpdateTaskStatusUsecase domain_usecase.UpdateTaskStatusUsecaseInterface
}

type Repositories struct {
	TaskRepository domain_repository.TaskRepositoryInterface
	UserRepository domain_repository.UserRepositoryInterface
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

	kafkaProducer := kafka_pkg.NewProducer(configs.KafkaCfg.TaskStatusUpdatedTopic)

	srv := web.NewServer(ctx, usecases.CreateTaskUsecase, usecases.CreateUserUsecase, usecases.GetUserUsecase, usecases.AuthUsecase, usecases.ListTaskUsecase, usecases.UpdateTaskStatusUsecase, kafkaProducer)

	return srv
}

func NewUsecases(ctx context.Context, repositories Repositories) Usecases {
	createTaskUsecase := usecase.NewCreateTaskUsecase(repositories.TaskRepository)
	createUserUsecase := usecase.NewCreateUserUsecase(repositories.UserRepository)
	getUserUsecase := usecase.NewGetUserUsecase(repositories.UserRepository)
	authUsecase := usecase.NewAuthUsecase(repositories.UserRepository)
	listTaskUsecase := usecase.NewListTaskUsecase(repositories.TaskRepository, repositories.UserRepository)
	updateTaskStatusUsecase := usecase.NewUpdateTaskStatusUsecase(repositories.TaskRepository)

	return Usecases{
		CreateTaskUsecase:       createTaskUsecase,
		CreateUserUsecase:       createUserUsecase,
		GetUserUsecase:          getUserUsecase,
		AuthUsecase:             authUsecase,
		ListTaskUsecase:         listTaskUsecase,
		UpdateTaskStatusUsecase: updateTaskStatusUsecase,
	}
}

func NewRepositories(ctx context.Context, db *gorm.DB) Repositories {
	taskRepository := repository_task.NewTaskRepository(db)
	userRepository := repository_user.NewUserRepository(db)

	return Repositories{
		TaskRepository: taskRepository,
		UserRepository: userRepository,
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
	err := db.AutoMigrate(&repository_user.User{})
	if err != nil {
		log.Fatal("Could not run migrations':", err)
	}

	return err
}

func NewMigrations(db *gorm.DB) error {
	err := NewUserMigration(db)

	if err != nil {
		log.Fatal("Could not run user migrations':", err)
	}

	err = NewTaskMigration(db)

	if err != nil {
		log.Fatal("Could not run task migrations':", err)
	}

	return err
}
