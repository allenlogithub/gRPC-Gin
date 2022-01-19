package controllers

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

var (
	addr = flag.String("addr", "172.17.0.6:4040", "the address to connect to")
)

// regist an user
// check the redis whether the account exists, if true, asked for email confirmation, else
// check the user database whether the account exists, if true, notifying the user the account has been
// registered, else
// add this registry record into the redis (set TTL)
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

	connn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connn.Close()

	cli := proto.NewRegisterServiceClient(connn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s := proto.RegisterRequest{
		Account:         r.Account,
		Name:            r.Name,
		Password:        r.Password,
		ConfirmPassword: r.ConfirmPassword,
	}
	res, err := cli.SetRegister(ctx, &s)
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
