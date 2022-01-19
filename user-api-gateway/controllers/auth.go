package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	client "user-api-gateway/client"
	proto "user-api-gateway/proto"
)

type (
	login struct {
		Account  string `json:"Account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)

func (u UserController) Login(c *gin.Context) {
	var r login
	if err := c.BindJSON(&r); err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s := proto.LoginRequest{
		Account:  r.Account,
		Password: r.Password,
	}
	res, err := client.GetAuthCli().Login(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "LoginFailed",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
		"err":     nil,
	})

	return
}
