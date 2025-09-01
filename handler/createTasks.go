package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joaquimsnjunior/todo-api/repository"
	"github.com/joaquimsnjunior/todo-api/schemes"
)

// CreateTasks é responsável por criar uma nova tarefa
func CreateTasks(c *gin.Context) {
	var t schemes.Tasks
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := repository.CreateTasks(c.Request.Context(), &t); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar tarefa"})
		return
	}
	c.JSON(http.StatusCreated, t)
}
