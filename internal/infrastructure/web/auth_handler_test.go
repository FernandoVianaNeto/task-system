package web_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"task-system/internal/domain/dto"
	domain_response "task-system/internal/domain/response"
	domain_usecase "task-system/internal/domain/usecase"
	"task-system/internal/infrastructure/web"
	mock_usecase "task-system/test/mocks/domain_usecase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUsecase := mock_usecase.NewMockAuthUsecaseInterface(ctrl)

	server := &web.Server{
		Usecases: domain_usecase.Usecases{
			AuthUsecase: mockAuthUsecase,
		},
	}

	mockAuthUsecase.EXPECT().
		Execute(gomock.Any(), gomock.Any()).
		Return(domain_response.AuthResponse{Token: "mocked-token"}, nil)

	requestBody, _ := json.Marshal(dto.AuthDto{
		Email:    "test@email.com",
		Password: "password123",
	})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequestWithContext(context.Background(), "POST", "/auth", bytes.NewBuffer(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	server.AuthHandler(ctx)

	assert.Equal(t, http.StatusOK, w.Code)

	var response domain_response.AuthResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, domain_response.AuthResponse{Token: "mocked-token"}, response)
}

func TestAuthHandler_InvalidBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUsecase := mock_usecase.NewMockAuthUsecaseInterface(ctrl)
	server := &web.Server{
		Usecases: domain_usecase.Usecases{
			AuthUsecase: mockAuthUsecase,
		},
	}

	requestBody := []byte("invalid-json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequestWithContext(context.Background(), "POST", "/auth", bytes.NewBuffer(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	server.AuthHandler(ctx)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Contains(t, w.Body.String(), "Body not compatible with expected")
}

func TestAuthHandler_AuthFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUsecase := mock_usecase.NewMockAuthUsecaseInterface(ctrl)
	server := &web.Server{
		Usecases: domain_usecase.Usecases{
			AuthUsecase: mockAuthUsecase,
		},
	}

	mockAuthUsecase.EXPECT().
		Execute(gomock.Any(), gomock.Any()).
		Return(domain_response.AuthResponse{}, errors.New("invalid credentials"))

	requestBody, _ := json.Marshal(dto.AuthDto{
		Email:    "test@email.com",
		Password: "wrongpassword",
	})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequestWithContext(context.Background(), "POST", "/auth", bytes.NewBuffer(requestBody))
	ctx.Request.Header.Set("Content-Type", "application/json")

	server.AuthHandler(ctx)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Contains(t, w.Body.String(), "invalid credentials")
}
