package usecase

import (
	"context"
	"task-system/internal/domain/dto"
	domain_repository "task-system/internal/domain/repository"
	domain_usecase "task-system/internal/domain/usecase"
)

type DeleteTaskUsecase struct {
	TaskRepository domain_repository.TaskRepositoryInterface
}

func NewDeleteTaskUsecase(
	repository domain_repository.TaskRepositoryInterface,
) domain_usecase.DeleteTaskUsecaseInterface {
	return &DeleteTaskUsecase{
		TaskRepository: repository,
	}
}

func (d *DeleteTaskUsecase) Execute(ctx context.Context, input dto.DeleteTaskDto) error {
	err := d.TaskRepository.DeleteTaskByUuid(ctx, input)

	return err
}
