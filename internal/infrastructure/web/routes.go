package web

import (
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine, server *Server) *gin.Engine {
	{
		task := engine.Group("/task")
		{
			task.POST("/", server.CreateTaskHandler)
		}
	}

	{
		user := engine.Group("/user")
		{
			user.POST("/", server.CreateUserHandler)
			user.GET("/:uuid", server.GetUserHandler)
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
