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
			postGroup.DELETE("/article", middlewares.JWTValidationMiddleware(), post.DelArticle)
			postGroup.POST("/articlecomment", middlewares.JWTValidationMiddleware(), post.PostArticleComment)
			postGroup.DELETE("/articlecomment", middlewares.JWTValidationMiddleware(), post.DelArticleComment)
			postGroup.GET("/personalarticle", middlewares.JWTValidationMiddleware(), post.GetPersonalArticle)
			// postGroup.GET("/friendarticle", middlewares.JWTValidationMiddleware(), post.GetFriendArticle)
		}
	}

	return router
}
