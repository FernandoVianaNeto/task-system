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
}

func NewListTaskUsecase(
	repository domain_repository.TaskRepositoryInterface,
) domain_usecase.ListTaskUsecaseInterface {
	return &ListTaskUsecase{
		TaskRepository: repository,
	}
}

func (l *ListTaskUsecase) Execute(ctx context.Context, input dto.ListTaskDto) ([]*entities.Task, error) {
	entity, err := l.TaskRepository.ListTask(ctx, input)

	return entity, err
}
