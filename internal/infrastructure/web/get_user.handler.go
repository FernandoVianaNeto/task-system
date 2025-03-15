package web

import (
	"net/http"
	"task-system/internal/domain/dto"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetUserHandler(ctx *gin.Context) {
	var req dto.GetUserDto

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "missing user uuid on params"})
		return
	}

	response, err := s.GetUserUsecase.Execute(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
