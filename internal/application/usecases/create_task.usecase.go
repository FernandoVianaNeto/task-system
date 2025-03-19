package usecase

import (
	"context"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
	domain_repository "task-system/internal/domain/repository"
	domain_usecase "task-system/internal/domain/usecase"
)

type CreateTaskUsecase struct {
	TaskRepository domain_repository.TaskRepositoryInterface
}

func NewCreateTaskUsecase(
	repository domain_repository.TaskRepositoryInterface,
) domain_usecase.CreateTaskUseCaseInterface {
	return &CreateTaskUsecase{
		TaskRepository: repository,
	}
}

func (c *CreateTaskUsecase) Execute(ctx context.Context, input dto.CreateTaskDto) error {
	taskEntity := entities.NewTask(
		input.UserUuid,
		input.Title,
		input.Summary,
	)

	err := c.TaskRepository.CreateTask(ctx, *taskEntity)

	return err
}
