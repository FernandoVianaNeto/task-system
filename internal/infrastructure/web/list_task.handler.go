package web

import (
	"net/http"
	"task-system/internal/domain/dto"
	infra_request "task-system/internal/infrastructure/web/request"

	"github.com/gin-gonic/gin"
)

func (s *Server) ListTasksHandler(ctx *gin.Context) {
	var req infra_request.ListTaskRequest
	var listTaskDto dto.ListTaskDto

	ctx.ShouldBindQuery(&req)

	if req.Owner != "" {
		listTaskDto.Owner = &req.Owner
	}

	if req.Uuid != "" {
		listTaskDto.Uuid = &req.Uuid
	}

	response, err := s.ListTaskUsecase.Execute(ctx, listTaskDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
