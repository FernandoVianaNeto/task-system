package domain_repository

import (
	"context"
	"task-system/internal/domain/entities"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type TaskRepositoryInterface interface {
	GetTaskByUser(ctx context.Context, taskUuid string, userUuid string) (*entities.Task, error)
	CreateTask(ctx context.Context, userUuid string, input entities.Task) error
	UpdateTaskByUser(ctx context.Context, userUuid string, input entities.Task) error
}
