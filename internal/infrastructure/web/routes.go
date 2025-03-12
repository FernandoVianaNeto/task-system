package web

import (
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine, server *Server) *gin.Engine {
	{
		// task := engine.Group("/task")
		// {
		// 	task.POST("/suggested-truckers-extension", server.ConsumePlanHandler)
		// }
	}

	return engine
}
