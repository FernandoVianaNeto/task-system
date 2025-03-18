package domain_repository

import (
	"context"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type TaskRepositoryInterface interface {
	CreateTask(ctx context.Context, input entities.Task) error
	ListTask(ctx context.Context, input dto.ListTaskDto) ([]*entities.Task, error)
	UpdateTaskStatus(ctx context.Context, input dto.UpdateTaskStatusDto) error
}
