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
		ArticleId int64  `json:"ArticleId" binding:"required"`
		Content   string `json:"Content" binding:"required"`
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
	s := proto.AddArticleRequest{
		UserId:     c.MustGet("UserId").(int64),
		Content:    r.Content,
		Visibility: r.Visibility,
	}
	res, err := client.GetPostArticleCli().AddArticle(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
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

func (p PostController) PostArticleComment(c *gin.Context) {
	var r postArticleComment
	if err := c.BindJSON(&r); err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s := proto.AddArticleCommentRequest{
		UserId:    c.MustGet("UserId").(int64),
		ArticleId: r.ArticleId,
		Content:   r.Content,
	}
	res, err := client.GetPostArticleCli().AddArticleComment(ctx, &s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "BadRequest",
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
