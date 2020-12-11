package main

import (
	handler "github.com/api/handler/rest"
	"github.com/api/infra/persistence"
	"github.com/api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	todoPersistence := persistence.NewTodoPersistence()
	todoUseCase := usecase.NewTodoUseCase(todoPersistence)
	todoHandler := handler.NewTodokHandler(todoUseCase)

	engine := gin.Default()
	engine.GET("/", todoHandler.Index)
	engine.Run(":4000")
}
