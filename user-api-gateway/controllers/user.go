package controllers

import (
// "net/http"

// "github.com/gin-gonic/gin"
// "github.com/google/uuid"
)

type (
	UserController struct{}

	login struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)

// // var userModel = new(models.User)

// // each divece will be given an uuid for recognization
// // all the devices using different token to access
// func (u UserController) Login(c *gin.Context) {
// 	var l login
// 	err := c.BindJSON(&l)
// 	if err != nil {
// 		c.JSON(c.Writer.Status(), gin.H{
// 			"message": "bad request",
// 			"err":     err,
// 		})
// 		return
// 	}
// 	// query the hash of the user from db
// 	// tbc...

// 	if !crypto.ComparePassword(hash, l.Password) {
// 		c.JSON(400, gin.H{
// 			"message": "bad request",
// 			"err":     nil,
// 		})
// 		return
// 	}
// 	// store login ingo into redis
// 	accessToken, device := token, uuid.New().String()
// 	// tbc...

// 	// return login info
// 	// the middleware will check the accessToken and device if required
// 	c.JSON(200, gin.H{
// 		"message": "bad request",
// 		"data": gin.H{
// 			"accessToken": accessToken,
// 			"device":      device,
// 		},
// 		"err": nil,
// 	})
// 	return
// }

// func (u UserController) Logout(c *gin.Context) {
// 	return
// }

// func (u UserController) LogoutAllDevices(c *gin.Context) {
// 	return
// }
