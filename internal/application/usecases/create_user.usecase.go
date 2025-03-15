package usecase

import (
	"context"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
	domain_repository "task-system/internal/domain/repository"
	domain_usecase "task-system/internal/domain/usecase"
)

type CreateUserUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
}

func NewCreateUserUsecase(
	repository domain_repository.UserRepositoryInterface,
) domain_usecase.CreateUserUsecaseInterface {
	return &CreateUserUsecase{
		UserRepository: repository,
	}
}

func (c *CreateUserUsecase) Execute(ctx context.Context, input dto.CreateUserDto) error {
	userEntity := entities.NewUser(
		input.Role,
		input.Name,
	)

	err := c.UserRepository.CreateUser(ctx, *userEntity)

	return err
}
