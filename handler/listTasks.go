package handler

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/joaquimsnjunior/todo-api/repository"
)

func GetTasksByID(c *gin.Context) {
	var id int
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da tarefa inválido"})
		return
	}
	task, err := repository.GetTasksByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar tarefa"})
		return
	}
	if task == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
		return
	}
	c.JSON(http.StatusOK, task)
}

