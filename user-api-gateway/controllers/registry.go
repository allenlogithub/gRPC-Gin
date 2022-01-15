package controllers

import (
	"context"
	"flag"
	"fmt"
	"log"
	// "net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	hs "user-api-gateway/crypto"
	proto "user-api-gateway/proto"
)

type (
	registry struct {
		Account         string `json:"account" binding:"required"`
		Name            string `json:"name" binding:"required"`
		Email           string `json:"email" binding:"required"`
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
func (u UserController) Registry(c *gin.Context) {
	var r registry
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "bad request",
			"err":     err,
		})
		return
	}
	hash, err := hs.HashAndSalt(r.Password)
	fmt.Println(hash)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "registry failed",
			"err":     err,
		})
		return
	}
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": gin.H{
	// 		"account":         r.Account,
	// 		"name":            r.Name,
	// 		"email":           r.Email,
	// 		"password":        r.Password,
	// 		"confirmpassword": r.ConfirmPassword,
	// 	},
	// })
	var s proto.RegistryRequest
	s.Account = r.Account
	s.Name = r.Name
	s.Email = r.Email
	s.Password = r.Password
	s.ConfirmPassword = r.ConfirmPassword

	connn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connn.Close()
	cli := proto.NewRegistryServiceClient(connn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	res, err := cli.SetRegistry(ctx, &s)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println(res)
	return
}
