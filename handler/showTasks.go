package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	// Logic to retrieve tasks
	c.JSON(http.StatusOK, gin.H{"tasks": []string{"Task 1", "Task 2"}})
}
