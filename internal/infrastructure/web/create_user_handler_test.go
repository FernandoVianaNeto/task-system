package web_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task-system/internal/domain/dto"
	domain_usecase "task-system/internal/domain/usecase"
	"task-system/internal/infrastructure/web"
	mock_usecase "task-system/test/mocks/domain_usecase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateUserUsecase := mock_usecase.NewMockCreateUserUsecaseInterface(ctrl)

	server := &web.Server{
		Usecases: domain_usecase.Usecases{
			CreateUserUsecase: mockCreateUserUsecase,
		},
	}

	mockCreateUserUsecase.EXPECT().
		Execute(gomock.Any(), gomock.Any()).
		Return(nil)

	requestBody, _ := json.Marshal(dto.CreateUserDto{
		Role:     "admin",
		Name:     "Test User",
		Email:    "test@email.com",
		Password: "securepassword",
	})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequestWithContext(context.Background(), "POST", "/user", bytes.NewBuffer(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	server.CreateUserHandler(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateUserHandler_InvalidBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateUserUsecase := mock_usecase.NewMockCreateUserUsecaseInterface(ctrl)

	server := &web.Server{
		Usecases: domain_usecase.Usecases{
			CreateUserUsecase: mockCreateUserUsecase,
		},
	}

	requestBody := []byte("invalid-json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequestWithContext(context.Background(), "POST", "/user", bytes.NewBuffer(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	server.CreateUserHandler(ctx)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Contains(t, w.Body.String(), "Body not compatible with expected")
}

func TestCreateUserHandler_UsecaseFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreateUserUsecase := mock_usecase.NewMockCreateUserUsecaseInterface(ctrl)

	server := &web.Server{
		Usecases: domain_usecase.Usecases{
			CreateUserUsecase: mockCreateUserUsecase,
		},
	}

	mockCreateUserUsecase.EXPECT().
		Execute(gomock.Any(), gomock.Any()).
		Return(errors.New("email already in use"))

	requestBody, _ := json.Marshal(dto.CreateUserDto{
		Role:     "user",
		Name:     "Test User",
		Email:    "test@email.com",
		Password: "securepassword",
	})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequestWithContext(context.Background(), "POST", "/user", bytes.NewBuffer(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	server.CreateUserHandler(ctx)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Contains(t, w.Body.String(), "email already in use")
}
