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
	sendFriendRequest struct {
		ReceiverUserId int64 `json:"ReceiverUserId" binding:"required"`
	}

	acceptFriendRequest struct {
		RequestorUserId int64 `json:"RequestorUserId" binding:"required"`
	}
)

func (u UserController) SendFriendRequest(c *gin.Context) {
	var r sendFriendRequest
	if err := c.BindJSON(&r); err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
		})
		return
	}

	if c.MustGet("UserId").(int64) == r.ReceiverUserId {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "SendFriendRequestFailed",
			"err":     "Requestor==Receiver",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s := proto.SendFriendRequestRequest{
		RequestorUserId: c.MustGet("UserId").(int64),
		ReceiverUserId:  r.ReceiverUserId,
	}
	res, err := client.GetUserPostCli().SendFriendRequest(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "SendFriendRequestFailed",
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

func (u UserController) AcceptFriendRequest(c *gin.Context) {
	var r acceptFriendRequest
	if err := c.BindJSON(&r); err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s := proto.AcceptFriendRequestRequest{
		RequestorUserId: r.RequestorUserId,
		ReceiverUserId:  c.MustGet("UserId").(int64),
	}
	res, err := client.GetUserPostCli().AcceptFriendRequest(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "AcceptFriendRequestFailed",
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

func (u UserController) GetFriendList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s := proto.GetFriendListRequest{
		UserId: c.MustGet("UserId").(int64),
	}
	res, err := client.GetUserGetFriendCli().GetFriendList(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "gRPC.GetFriendListFailed",
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

func (u UserController) SearchUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	searchString := c.Query("SearchString")
	if len(searchString) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "InvalidSearchString",
			"err":     "InvalidSearchString",
		})
		return
	}
	s := proto.SearchUserRequest{
		UserId:       c.MustGet("UserId").(int64),
		SearchString: searchString,
	}
	res, err := client.GetUserGetFriendCli().SearchUser(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "gRPC.SearchUserFailed",
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

func (u UserController) GetFriendRequestList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s := proto.GetFriendRequestListRequest{
		UserId: c.MustGet("UserId").(int64),
	}
	res, err := client.GetUserGetFriendCli().GetFriendRequestList(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "gRPC.GetFriendRequestListFailed",
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
