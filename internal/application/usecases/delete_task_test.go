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

func TestDeleteTaskUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepository := mock_repository.NewMockTaskRepositoryInterface(ctrl)
	deleteTaskUsecase := usecase.NewDeleteTaskUsecase(mockTaskRepository)

	t.Run("should return error if repository fails to delete task", func(t *testing.T) {
		mockTaskRepository.EXPECT().
			DeleteTaskByUuid(gomock.Any(), gomock.Any()).
			Return(errors.New("failed to delete task"))

		err := deleteTaskUsecase.Execute(context.Background(), dto.DeleteTaskDto{
			Uuid: "123e4567-e89b-12d3-a456-426614174000",
		})

		assert.Error(t, err)
		assert.Equal(t, "failed to delete task", err.Error())
	})

	t.Run("should delete task successfully", func(t *testing.T) {
		mockTaskRepository.EXPECT().
			DeleteTaskByUuid(gomock.Any(), gomock.Any()).
			Return(nil)

		err := deleteTaskUsecase.Execute(context.Background(), dto.DeleteTaskDto{
			Uuid: "123e4567-e89b-12d3-a456-426614174000",
		})

		assert.NoError(t, err)
	})
}
