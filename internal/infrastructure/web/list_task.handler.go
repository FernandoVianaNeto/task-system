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

func (s *Server) ListTasksHandler(ctx *gin.Context) {
	var req infra_request.ListTaskRequest
	var listTaskDto dto.ListTaskDto

	ctx.ShouldBindQuery(&req)

	token := ctx.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", 1)
	parsedToken, err := jwt.ParseWithClaims(token, &middleware.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.ApplicationCfg.PasswordSecretHash), nil
	})

	if req.Owner != "" {
		listTaskDto.Owner = &req.Owner
	}

	if req.Uuid != "" {
		listTaskDto.Uuid = &req.Uuid
	}

	if claims, ok := parsedToken.Claims.(*middleware.JwtClaims); ok {
		listTaskDto.Role = claims.Role
		listTaskDto.UserUuid = claims.Uuid
	}

	response, err := s.ListTaskUsecase.Execute(ctx, listTaskDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
