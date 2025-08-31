package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTasks(c *gin.Context) {
	// Logic to create a new task
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}
