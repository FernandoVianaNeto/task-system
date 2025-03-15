package domain_repository

import (
	"context"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, input entities.User) error
	GetUserByUuid(ctx context.Context, input dto.GetUserByUuidDto) (*entities.User, error)
	GetUserByEmail(ctx context.Context, input dto.GetUserByEmailDto) (*entities.User, error)
}
