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

	// post
	postPersistence := persistence.NewPostPersistence()
	potsUseCase := usecase.NewPostCase(postPersistence)
	postHandler := handler.NewPostHandler(potsUseCase)

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

				// user
				v1.DELETE("/user", userHandler.DeleteUser)

				// refreshToken
				v1.PATCH("/refresh_token", jwtHandler.RefreshToken)

				// posts
				v1.GET("/posts", postHandler.GetPostAll)

				self := v1.Group("self")
				{
					// user
					self.GET("/user", userHandler.GetCurrentUser)
					self.PATCH("/user", userHandler.UpdateUser)

					// post
					self.POST("/post", postHandler.AddPost)
					self.GET("/posts", postHandler.GetCurrentPosts)
					self.PATCH("/post/:id", postHandler.UpdatePost)
					self.DELETE("/post/:id", postHandler.DeletePost)
				}
			}
		}
	}

	return engine
}
