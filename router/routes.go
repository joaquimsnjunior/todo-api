package router

import (
	"github.com/gin-gonic/gin"
	"github.com/joaquimsnjunior/todo-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	// initialize your routes here
	handler.InitializeHandlers()
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/tasks", handler.GetTasks)
		v1.POST("/tasks", handler.CreateTasks)
		v1.GET("/tasks/:id", handler.GetTasksByID)
		v1.PUT("/tasks/:id", handler.UpdateTasks)
		v1.DELETE("/tasks/:id", handler.DeleteTasks)
	}
}
