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

	return engine
}
