package usecase_test

import (
	"context"
	"errors"
	"testing"

	configs "task-system/cmd/config"
	usecase "task-system/internal/application/usecases"
	"task-system/internal/domain/dto"
	"task-system/internal/domain/entities"
	mock_repository "task-system/test/mocks/domain_repository"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func ConfigureTestConfig() {
	configs.ApplicationCfg = &configs.ApplicationConfig{}
}

func TestAuthUsecase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repository.NewMockUserRepositoryInterface(ctrl)

	authUsecase := usecase.NewAuthUsecase(mockUserRepository)

	ConfigureTestConfig()

	t.Run("should return error if user not found", func(t *testing.T) {
		mockUserRepository.EXPECT().
			GetUserByEmail(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("user not found"))

		_, err := authUsecase.Execute(context.Background(), dto.AuthDto{
			Email:    "invalid@email.com",
			Password: "password123",
		})

		assert.Equal(t, "user not found", err.Error())
	})

	t.Run("should return error if password is incorrect", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

		mockUserRepository.EXPECT().
			GetUserByEmail(gomock.Any(), gomock.Any()).
			Return(&entities.User{
				Email:    "valid@email.com",
				Password: string(hashedPassword),
				Role:     "admin",
				Uuid:     "12345",
			}, nil)

		_, err := authUsecase.Execute(context.Background(), dto.AuthDto{
			Email:    "valid@email.com",
			Password: "wrongpassword",
		})

		assert.Error(t, err)
	})

	t.Run("should return token if authentication is successful", func(t *testing.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

		configs.ApplicationCfg.PasswordSecretHash = "secret_hash"

		mockUserRepository.EXPECT().
			GetUserByEmail(gomock.Any(), gomock.Any()).
			Return(&entities.User{
				Email:    "valid@email.com",
				Password: string(hashedPassword),
				Role:     "admin",
				Uuid:     "12345",
			}, nil)

		response, err := authUsecase.Execute(context.Background(), dto.AuthDto{
			Email:    "valid@email.com",
			Password: "password123",
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, response.Token)
	})
}
