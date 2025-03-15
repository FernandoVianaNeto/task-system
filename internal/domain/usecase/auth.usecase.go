package domain_usecase

import (
	"context"
	"task-system/internal/domain/dto"
	domain_response "task-system/internal/domain/response"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type AuthUsecaseInterface interface {
	Execute(ctx context.Context, input dto.AuthDto) (domain_response.AuthResponse, error)
}
