package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	client "post-api-gateway/client"
	proto "post-api-gateway/proto"
)

func (p PostController) GetPersonalArticle(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s := proto.GetPersonalArticleRequest{
		UserId: c.MustGet("UserId").(int64),
	}
	res, err := client.GetGetArticleCli().GetPersonalArticle(ctx, &s)
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

// func (p PostController) GetFriendArticle(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	s := proto.GetFriendArticleRequest{
// 		UserId: c.MustGet("UserId").(int64),
// 	}
// 	res, err := client.GetGetArticleCli().GetFriendArticle(ctx, &s)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "BadRequest",
// 			"err":     err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": res,
// 		"err":     nil,
// 	})

// 	return
// }
