package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter() *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		indexController(ctx)
	})

	router.POST("users", func(ctx *gin.Context) {
		CreateUserController(ctx)
	})

	router.GET("users", func(ctx *gin.Context) {
		GetAllUserController(ctx)
	})

	router.GET("users/:id", func(ctx *gin.Context) {
		GetUserByIdController(ctx)
	})

	router.PUT("users/:id", func(ctx *gin.Context) {
		UpdateUserByIdController(ctx)
	})

	router.DELETE("users/:id", func(ctx *gin.Context) {
		DeleteUserByIdController(ctx)
	})

	return router
}
