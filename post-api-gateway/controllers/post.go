package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	client "post-api-gateway/client"
	proto "post-api-gateway/proto"
)

type (
	PostController struct{}

	postArticle struct {
		// AccessToken string `json:"AccessToken" binding:"required"`
		Content    string `json:"Content" binding:"required"`
		Visibility string `json:"Visibility" binding:"required"`
	}

	postArticleAttachment struct {
		// AccessToken string
		FileStream string
		FileName   string
	}

	postComment struct {
		// AccessToken string `json:"AccessToken" binding:"required"`
		Content string `json:"Content" binding:"required"`
	}
)

func (p PostController) PostArticle(c *gin.Context) {
	var r postArticle
	if err := c.BindJSON(&r); err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
		})
		return
	}
	tk, err1 := c.Cookie("AccessToken")
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": nil,
			"err":     err1.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s := proto.PostPostRequest{
		AccessToken: tk,
		Content:     r.Content,
		Visibility:  r.Visibility,
	}
	res, err2 := client.GetPostPostCli().PostPost(ctx, &s)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
			"err":     err2.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": res,
		"err":     nil,
	})

	return
}
