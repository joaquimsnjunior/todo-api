package handler

import "github.com/gin-gonic/gin"

func InitializeHandlers() {
	r := gin.Default()
	r.GET("/tasks", GetTasks)
	r.POST("/tasks", CreateTasks)
	r.GET("/tasks/:id", GetTasksByID)
	r.PUT("/tasks/:id", UpdateTasks)
	r.DELETE("/tasks/:id", DeleteTasks)
}
