package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteTasks(c *gin.Context) {
	// Logic to delete a task
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
