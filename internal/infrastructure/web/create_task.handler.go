package web

import (
	"net/http"
	"task-system/internal/domain/dto"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateTaskHandler(ctx *gin.Context) {
	var req dto.CreateTaskDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Body not compatible with expected"})
		return
	}

	err := s.CreateTaskUsecase.Execute(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
