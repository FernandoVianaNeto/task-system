package usecase

import (
	"context"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
	domain_repository "task-system/internal/domain/repository"
	domain_usecase "task-system/internal/domain/usecase"
)

type ListTaskUsecase struct {
	TaskRepository domain_repository.TaskRepositoryInterface
	UserRepository domain_repository.UserRepositoryInterface
}

func NewListTaskUsecase(
	taskRepository domain_repository.TaskRepositoryInterface,
	userRepository domain_repository.UserRepositoryInterface,
) domain_usecase.ListTaskUsecaseInterface {
	return &ListTaskUsecase{
		TaskRepository: taskRepository,
		UserRepository: userRepository,
	}
}

func (l *ListTaskUsecase) Execute(ctx context.Context, input dto.ListTaskDto) ([]*entities.Task, error) {
	entity, err := l.TaskRepository.ListTask(ctx, input)

	return entity, err
}
