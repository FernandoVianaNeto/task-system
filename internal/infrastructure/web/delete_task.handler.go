package web

import (
	"net/http"
	"task-system/internal/domain/dto"
	infra_request "task-system/internal/infrastructure/web/request"

	"github.com/gin-gonic/gin"
)

func (s *Server) DeleteTaskHandler(ctx *gin.Context) {
	var req infra_request.DeleteTaskRequestParam

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "missing task uuid on params"})
		return
	}

	err := s.Usecases.DeleteTaskUsecase.Execute(ctx, dto.DeleteTaskDto{Uuid: req.Uuid})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
