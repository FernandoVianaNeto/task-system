package usecase

import (
	"context"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
	domain_repository "task-system/internal/domain/repository"
	domain_service "task-system/internal/domain/service"
	domain_usecase "task-system/internal/domain/usecase"
)

type CreateUserUsecase struct {
	UserRepository        domain_repository.UserRepositoryInterface
	PasswordHasherService domain_service.PasswordHasherServiceInterface
}

func NewCreateUserUsecase(
	repository domain_repository.UserRepositoryInterface,
	hasherService domain_service.PasswordHasherServiceInterface,
) domain_usecase.CreateUserUsecaseInterface {
	return &CreateUserUsecase{
		UserRepository:        repository,
		PasswordHasherService: hasherService,
	}
}

func (c *CreateUserUsecase) Execute(ctx context.Context, input dto.CreateUserDto) error {
	hashedPassword, err := c.PasswordHasherService.HashPassword(input.Password, 10)

	if err != nil {
		return err
	}

	userEntity := entities.NewUser(
		input.Role,
		input.Name,
		input.Email,
		string(hashedPassword),
	)

	err = c.UserRepository.CreateUser(ctx, *userEntity)

	return err
}
