package router

import (
	handler "github.com/api/handler/rest"
	"github.com/api/infra/persistence"
	"github.com/api/usecase"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	engine := gin.Default()

	// =====================================
	// 依存関係を注入
	// =====================================

	// example
	todoPersistence := persistence.NewTodoPersistence()
	todoUseCase := usecase.NewTodoUseCase(todoPersistence)
	todoHandler := handler.NewTodokHandler(todoUseCase)

	// =====================================
	// ルーティング
	// =====================================

	engine.GET("/example", todoHandler.Index)

	return engine
}
