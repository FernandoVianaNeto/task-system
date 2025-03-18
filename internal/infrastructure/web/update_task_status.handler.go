package web

import (
	"net/http"
	"strings"
	configs "task-system/cmd/config"
	"task-system/internal/domain/dto"
	"task-system/internal/infrastructure/web/middleware"
	infra_request "task-system/internal/infrastructure/web/request"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Server) UpdateTaskStatusHandler(ctx *gin.Context) {
	var req infra_request.UpdateTaskStatusRequest
	var updateTaskStatusDto dto.UpdateTaskStatusDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Body not compatible with expected"})
		return
	}

	token := ctx.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)
	parsedToken, err := jwt.ParseWithClaims(token, &middleware.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.ApplicationCfg.PasswordSecretHash), nil
	})

	if claims, ok := parsedToken.Claims.(*middleware.JwtClaims); ok {
		updateTaskStatusDto.UserUuid = claims.Uuid
	}

	updateTaskStatusDto.TaskStatus = req.NewStatus
	updateTaskStatusDto.TaskUuid = req.TaskUuid

	err = s.UpdateTaskStatusUsecase.Execute(ctx, updateTaskStatusDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
