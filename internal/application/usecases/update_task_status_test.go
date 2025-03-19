package usecase_test

import (
	"context"
	"errors"
	"testing"

	usecase "task-system/internal/application/usecases"
	"task-system/internal/domain/dto"
	mock_repository "task-system/test/mocks/domain_repository"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateTaskStatusUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepository := mock_repository.NewMockTaskRepositoryInterface(ctrl)
	updateTaskStatusUsecase := usecase.NewUpdateTaskStatusUsecase(mockTaskRepository)

	t.Run("should return error if repository fails to update task status", func(t *testing.T) {
		mockTaskRepository.EXPECT().
			UpdateTaskStatus(gomock.Any(), gomock.Any()).
			Return(errors.New("failed to update task status"))

		err := updateTaskStatusUsecase.Execute(context.Background(), dto.UpdateTaskStatusDto{
			TaskUuid:   "123e4567-e89b-12d3-a456-426614174000",
			UserUuid:   "user-uuid",
			TaskStatus: "completed",
		})

		assert.Error(t, err)
		assert.Equal(t, "failed to update task status", err.Error())
	})

	t.Run("should update task status successfully", func(t *testing.T) {
		mockTaskRepository.EXPECT().
			UpdateTaskStatus(gomock.Any(), gomock.Any()).
			Return(nil)

		err := updateTaskStatusUsecase.Execute(context.Background(), dto.UpdateTaskStatusDto{
			TaskUuid:   "123e4567-e89b-12d3-a456-426614174000",
			UserUuid:   "user-uuid",
			TaskStatus: "completed",
		})

		assert.NoError(t, err)
	})
}
