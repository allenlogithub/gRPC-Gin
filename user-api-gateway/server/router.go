package server

import (
	"github.com/gin-gonic/gin"

	"user-api-gateway/controllers"
	// "github.com/vsouza/go-gin-boilerplate/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// health := new(controllers.HealthController)

	// router.GET("/health", health.Status)
	// router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			// userGroup.GET("/greeting/:id", user.Login)
			userGroup.POST("/registry", user.Register)
		}
	}
	return router
}
