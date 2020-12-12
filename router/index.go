package router

import (
	handler "github.com/api/handler/rest"
	"github.com/api/infra/persistence"
	"github.com/api/middleware"
	"github.com/api/usecase"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	engine := gin.Default()

	engine.Use(middleware.SetDB)

	// =====================================
	// 依存関係を注入
	// =====================================

	// example
	todoPersistence := persistence.NewTodoPersistence()
	todoUseCase := usecase.NewTodoUseCase(todoPersistence)
	todoHandler := handler.NewTodokHandler(todoUseCase)

	// user
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserseCase(userPersistence)
	userHandler := handler.NewUserkHandler(userUseCase)

	// =====================================
	// ルーティング
	// =====================================

	v1 := engine.Group("v1")
	{
		// example
		v1.GET("/example", todoHandler.Index)

		// user
		v1.GET("/users", userHandler.GetUserAll)
		v1.POST("/user", userHandler.AddUser)
	}

	return engine
}
