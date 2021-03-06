package server

import (
	"github.com/gin-gonic/gin"

	controllers "user-api-gateway/controllers"
	middlewares "user-api-gateway/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.POST("/register", user.Register)
			userGroup.POST("/login", user.Login)
			userGroup.POST("/logout", middlewares.JWTValidationMiddleware(), user.Logout)
			userGroup.POST("/sendfriendrequest", middlewares.JWTValidationMiddleware(), user.SendFriendRequest)
			userGroup.POST("/acceptfriendrequest", middlewares.JWTValidationMiddleware(), user.AcceptFriendRequest)
			userGroup.DELETE("/rejectfriendrequest", middlewares.JWTValidationMiddleware(), user.RejectFriendRequest)
			userGroup.GET("/friendlist", middlewares.JWTValidationMiddleware(), user.GetFriendList)
			userGroup.GET("/searchuser", middlewares.JWTValidationMiddleware(), user.SearchUser)
			userGroup.GET("/friendrequestlist", middlewares.JWTValidationMiddleware(), user.GetFriendRequestList)
		}
	}

	return router
}
