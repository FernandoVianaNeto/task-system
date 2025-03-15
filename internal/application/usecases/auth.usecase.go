package usecase

import (
	"context"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
	domain_repository "task-system/internal/domain/repository"
	domain_usecase "task-system/internal/domain/usecase"
)

type AuthUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
}

func NewAuthUsecase(
	repository domain_repository.UserRepositoryInterface,
) domain_usecase.AuthUsecaseInterface {
	return &AuthUsecase{
		UserRepository: repository,
	}
}

func (c *AuthUsecase) Execute(ctx context.Context, input dto.AuthDto) error {
	taskEntity := entities.NewTask(
		input.UserUuid,
		input.Title,
		input.Summary,
	)

	err := c.UserRepository.GetUser(ctx, *taskEntity)

	return err
}
