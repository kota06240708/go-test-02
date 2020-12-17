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

	// refreshToken
	refreshTokenPersistence := persistence.NewRefreshTokenPersistence()
	refreshTokenUseCase := usecase.NewRefreshTokenCase(refreshTokenPersistence)

	// jwt
	jwtHandler := handler.NewJwtHandler(userUseCase, refreshTokenUseCase)

	// =====================================
	// ルーティング
	// =====================================

	api := engine.Group("api")
	{
		v1 := api.Group("v1")
		{
			// user
			v1.GET("/users", userHandler.GetUserAll)
			v1.POST("/signup", userHandler.AddUser)

			// example
			v1.GET("/example", todoHandler.Index)

			// jwt
			v1.POST("/login", jwtHandler.AuthMiddleware().LoginHandler)

			v1.Use(jwtHandler.AuthMiddleware().MiddlewareFunc())
			{
				self := v1.Group("self")
				{
					self.GET("/user", userHandler.GetCurrentUser)
				}

				// refreshToken
				v1.PATCH("/refresh_token", jwtHandler.RefreshToken)
			}
		}
	}

	return engine
}
