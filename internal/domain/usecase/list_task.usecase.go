package domain_usecase

import (
	"context"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type ListTaskUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ListTaskDto) ([]*entities.Task, error)
}
