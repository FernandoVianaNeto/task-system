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

func TestCreateTaskUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepository := mock_repository.NewMockTaskRepositoryInterface(ctrl)

	createTaskUsecase := usecase.NewCreateTaskUsecase(mockTaskRepository)

	t.Run("should return error if repository fails to create task", func(t *testing.T) {
		mockTaskRepository.EXPECT().
			CreateTask(gomock.Any(), gomock.Any()).
			Return(errors.New("failed to create task"))

		err := createTaskUsecase.Execute(context.Background(), dto.CreateTaskDto{
			UserUuid: "123e4567-e89b-12d3-a456-426614174000",
			Title:    "Test Task",
			Summary:  "This is a test task",
		})

		assert.Error(t, err)
		assert.Equal(t, "failed to create task", err.Error())
	})

	t.Run("should create task successfully", func(t *testing.T) {
		mockTaskRepository.EXPECT().
			CreateTask(gomock.Any(), gomock.Any()).
			Return(nil)

		err := createTaskUsecase.Execute(context.Background(), dto.CreateTaskDto{
			UserUuid: "123e4567-e89b-12d3-a456-426614174000",
			Title:    "Test Task",
			Summary:  "This is a test task",
		})

		assert.NoError(t, err)
	})
}
