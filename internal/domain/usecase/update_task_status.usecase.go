package domain_usecase

import (
	"context"
	"task-system/internal/domain/dto"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type UpdateTaskStatusUsecaseInterface interface {
	Execute(ctx context.Context, input dto.UpdateTaskStatusDto) error
}
