package server

import (
	"github.com/gin-gonic/gin"

	controllers "post-api-gateway/controllers"
	middlewares "post-api-gateway/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		postGroup := v1.Group("post")
		{
			post := new(controllers.PostController)
			postGroup.POST("/article", middlewares.JWTValidationMiddleware(), post.PostArticle)
			// postGroup.GET("/", middlewares.JWTValidationMiddleware(), post.GetArticles)
		}
	}

	return router
}
