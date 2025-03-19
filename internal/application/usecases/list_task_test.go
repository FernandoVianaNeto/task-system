package usecase_test

import (
	"context"
	"errors"
	"testing"

	usecase "task-system/internal/application/usecases"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
	mock_repository "task-system/test/mocks/domain_repository"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestListTaskUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepository := mock_repository.NewMockTaskRepositoryInterface(ctrl)
	mockUserRepository := mock_repository.NewMockUserRepositoryInterface(ctrl)

	listTaskUsecase := usecase.NewListTaskUsecase(mockTaskRepository, mockUserRepository)

	t.Run("should return error if repository fails to list tasks", func(t *testing.T) {
		mockTaskRepository.EXPECT().
			ListTask(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("failed to list tasks"))

		tasks, err := listTaskUsecase.Execute(context.Background(), dto.ListTaskDto{
			UserUuid: "123e4567-e89b-12d3-a456-426614174000",
		})

		assert.Nil(t, tasks)
		assert.Error(t, err)
		assert.Equal(t, "failed to list tasks", err.Error())
	})

	t.Run("should return list of tasks successfully", func(t *testing.T) {
		expectedTasks := []*entities.Task{
			{
				Uuid:    "task-1",
				Title:   "Task 1",
				Summary: "First test task",
				Owner:   "123e4567-e89b-12d3-a456-426614174001",
				Status:  "active",
			},
			{
				Uuid:    "task-2",
				Title:   "Task 2",
				Summary: "Second test task",
				Owner:   "123e4567-e89b-12d3-a456-426614174000",
				Status:  "active",
			},
		}

		mockTaskRepository.EXPECT().
			ListTask(gomock.Any(), gomock.Any()).
			Return(expectedTasks, nil)

		tasks, err := listTaskUsecase.Execute(context.Background(), dto.ListTaskDto{
			UserUuid: "123e4567-e89b-12d3-a456-426614174000",
		})

		assert.NoError(t, err)
		assert.NotNil(t, tasks)
		assert.Equal(t, expectedTasks, tasks)
	})
}
