package usecase

import (
	"context"
	"task-system/internal/domain/dto"
	domain_repository "task-system/internal/domain/repository"
	domain_usecase "task-system/internal/domain/usecase"
)

type UpdateTaskStatusUsecase struct {
	TaskRepository domain_repository.TaskRepositoryInterface
}

func NewUpdateTaskStatusUsecase(
	taskRepository domain_repository.TaskRepositoryInterface,
) domain_usecase.UpdateTaskStatusUsecaseInterface {
	return &UpdateTaskStatusUsecase{
		TaskRepository: taskRepository,
	}
}

func (u *UpdateTaskStatusUsecase) Execute(ctx context.Context, input dto.UpdateTaskStatusDto) error {
	err := u.TaskRepository.UpdateTaskStatus(ctx, input)

	return err
}
