package usecase

import (
	"context"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
	domain_repository "task-system/internal/domain/repository"
	domain_usecase "task-system/internal/domain/usecase"
)

type GetUserUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
}

func NewGetUserUsecase(
	repository domain_repository.UserRepositoryInterface,
) domain_usecase.GetUserUsecaseInterface {
	return &GetUserUsecase{
		UserRepository: repository,
	}
}

func (c *GetUserUsecase) Execute(ctx context.Context, input dto.GetUserByUuidDto) (*entities.User, error) {
	entity, err := c.UserRepository.GetUserByUuid(ctx, input)

	return entity, err
}
