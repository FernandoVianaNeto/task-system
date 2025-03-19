package web

import (
	"task-system/internal/infrastructure/web/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine, server *Server) *gin.Engine {

	{
		task := engine.Group("/task")
		{
			task.POST("/", middleware.JwtAuthMiddleware(), server.CreateTaskHandler)
			task.GET("/", middleware.JwtAuthMiddleware(), server.ListTasksHandler)
			task.PUT("/", middleware.JwtAuthMiddleware(), server.UpdateTaskStatusHandler)
			task.DELETE("/:uuid", middleware.JwtAdminAuthMiddleware(), server.DeleteTaskHandler)
		}
	}

	{
		user := engine.Group("/user")
		{
			user.POST("/", server.CreateUserHandler)
		}
	}

	{
		auth := engine.Group("/auth")
		{
			auth.POST("/", server.AuthHandler)
		}
	}

	return engine
}
