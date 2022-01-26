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
	delArticle struct {
		ArticleId int64 `json:"ArticleId" binding:"required"`
	}

	delArticleComment struct {
		ArticleCommentId int64 `json:"ArticleCommentId" binding:"required"`
	}
)

func (p PostController) DelArticle(c *gin.Context) {
	var r delArticle
	if err := c.BindJSON(&r); err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s := proto.DelArticleRequest{
		UserId:    c.MustGet("UserId").(int64),
		ArticleId: r.ArticleId,
	}
	res, err := client.GetPostArticleCli().DelArticle(ctx, &s)
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

func (p PostController) DelArticleComment(c *gin.Context) {
	var r delArticleComment
	if err := c.BindJSON(&r); err != nil {
		c.JSON(c.Writer.Status(), gin.H{
			"message": "BadRequest",
			"err":     err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s := proto.DelArticleCommentRequest{
		UserId:           c.MustGet("UserId").(int64),
		ArticleCommentId: r.ArticleCommentId,
	}
	res, err := client.GetPostArticleCli().DelArticleComment(ctx, &s)
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
