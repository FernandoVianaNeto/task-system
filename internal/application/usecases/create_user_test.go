package usecase_test

import (
	"context"
	"errors"
	"testing"

	usecase "task-system/internal/application/usecases"
	"task-system/internal/domain/dto"
	mock_repository "task-system/test/mocks/domain_repository"
	domain_service "task-system/test/mocks/domain_service"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repository.NewMockUserRepositoryInterface(ctrl)
	mockPasswordHasher := domain_service.NewMockPasswordHasherServiceInterface(ctrl)

	createUserUsecase := usecase.NewCreateUserUsecase(mockUserRepository, mockPasswordHasher)

	t.Run("should return error if hashing password fails", func(t *testing.T) {
		mockPasswordHasher.EXPECT().
			HashPassword(gomock.Any(), 10).
			Return(nil, errors.New("bcrypt error"))

		err := createUserUsecase.Execute(context.Background(), dto.CreateUserDto{
			Role:     "user",
			Name:     "Test User",
			Email:    "test@email.com",
			Password: "invalidpassword",
		})

		assert.Error(t, err)
		assert.Equal(t, "bcrypt error", err.Error())
	})

	t.Run("should return error if repository fails to create user", func(t *testing.T) {
		mockPasswordHasher.EXPECT().
			HashPassword(gomock.Any(), 10).
			Return([]byte("hashedpassword"), nil)

		mockUserRepository.EXPECT().
			CreateUser(gomock.Any(), gomock.Any()).
			Return(errors.New("failed to create user"))

		err := createUserUsecase.Execute(context.Background(), dto.CreateUserDto{
			Role:     "user",
			Name:     "Test User",
			Email:    "test@email.com",
			Password: "password123",
		})

		assert.Error(t, err)
		assert.Equal(t, "failed to create user", err.Error())
	})

	t.Run("should create user successfully", func(t *testing.T) {
		mockPasswordHasher.EXPECT().
			HashPassword(gomock.Any(), 10).
			Return([]byte("hashedpassword"), nil)

		mockUserRepository.EXPECT().
			CreateUser(gomock.Any(), gomock.Any()).
			Return(nil)

		err := createUserUsecase.Execute(context.Background(), dto.CreateUserDto{
			Role:     "user",
			Name:     "Test User",
			Email:    "test@email.com",
			Password: "password123",
		})

		assert.NoError(t, err)
	})
}
