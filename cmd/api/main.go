package main

import (
	"github.com/joaquimsnjunior/todo-api/config"
	"github.com/joaquimsnjunior/todo-api/router"
)

// test/start server
func main() {
	config.GetDB() // conecta ao banco de dados

	router.InitializeRouter()
}
