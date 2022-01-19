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
	register struct {
		Name            string `json:"name" binding:"required"`
		Account         string `json:"account" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}
)

// regist an user
func (u UserController) Register(c *gin.Context) {
	var r register
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s := proto.RegisterRequest{
		Account:         r.Account,
		Name:            r.Name,
		Password:        r.Password,
		ConfirmPassword: r.ConfirmPassword,
	}
	res, err := client.GetRegisterCli().SetRegister(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "RegisterFailed",
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
