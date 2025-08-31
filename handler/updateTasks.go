package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTasks(c *gin.Context) {
	// Logic to update a task
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}
