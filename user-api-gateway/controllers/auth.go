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

// If first login on the device (AccessToken not in redis),
// ask user sending (account, password) to login, else
// update the AccessToken within redis and cookie
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

	// type http.Cookie struct:
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	// gin.SetCookie:
	// name, value string, maxAge int, path, domain string, secure, httpOnly bool
	// use http -> secure=false; gin.Default().RunTLS() for https
	colonIdx := len(c.Request.Host) - 1
	for i := len(c.Request.Host) - 1; i >= 0; i-- {
		if c.Request.Host[i] == byte(58) {
			colonIdx = i
			break
		}
	}
	domain := c.Request.Host[:colonIdx]
	c.SetCookie("AccessToken", res.AccessToken, 0, "/", domain, false, false)
	c.JSON(http.StatusOK, gin.H{
		"message": "login Successfully",
		"err":     nil,
	})

	return
}

func (u UserController) Logout(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tk, err := c.Cookie("AccessToken")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": nil,
			"err":     err.Error(),
		})
		return
	}
	s := proto.LogoutRequest{
		AccessToken: tk,
	}
	res, err := client.GetAuthCli().Logout(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": res,
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
