package controllers

import (
	"context"
	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	client "post-api-gateway/client"
	proto "post-api-gateway/proto"
)

type (
	PostController struct{}

	postArticle struct {
		Content    string `json:"Content" binding:"required"`
		Visibility string `json:"Visibility" binding:"required"`
	}

	postArticleAttachment struct {
		FileStream string
		FileName   string
	}

	postArticleComment struct {
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s := proto.PostPostRequest{
		UserId:     c.MustGet("UserId").(int64),
		Content:    r.Content,
		Visibility: r.Visibility,
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
