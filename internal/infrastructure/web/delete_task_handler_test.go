package web_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	domain_usecase "task-system/internal/domain/usecase"
	"task-system/internal/infrastructure/web"
	mock_usecase "task-system/test/mocks/domain_usecase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDeleteTaskHandler_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeleteTaskUsecase := mock_usecase.NewMockDeleteTaskUsecaseInterface(ctrl)

	server := &web.Server{
		Usecases: domain_usecase.Usecases{
			DeleteTaskUsecase: mockDeleteTaskUsecase,
		},
	}

	mockDeleteTaskUsecase.EXPECT().
		Execute(gomock.Any(), gomock.Any()).
		Return(nil)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{{Key: "uuid", Value: "task-123"}}
	ctx.Request, _ = http.NewRequestWithContext(context.Background(), "DELETE", "/task/task-123", nil)

	server.DeleteTaskHandler(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTaskHandler_UsecaseFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDeleteTaskUsecase := mock_usecase.NewMockDeleteTaskUsecaseInterface(ctrl)

	server := &web.Server{
		Usecases: domain_usecase.Usecases{
			DeleteTaskUsecase: mockDeleteTaskUsecase,
		},
	}

	mockDeleteTaskUsecase.EXPECT().
		Execute(gomock.Any(), gomock.Any()).
		Return(errors.New("task not found"))

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = gin.Params{{Key: "uuid", Value: "task-123"}}
	ctx.Request, _ = http.NewRequestWithContext(context.Background(), "DELETE", "/task/task-123", nil)

	server.DeleteTaskHandler(ctx)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Contains(t, w.Body.String(), "task not found")
}
