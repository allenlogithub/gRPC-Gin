package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	client "post-api-gateway/client"
	proto "post-api-gateway/proto"
)

func JWTValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tk, err := c.Cookie("AccessToken")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		s := proto.JWTValidationRequest{
			AccessToken: tk,
		}
		res, err := client.GetAuthCli().ValidateJWT(ctx, &s)
		if err != nil || !res.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "AuthFailed"})
			return
		}
		c.Set("UserId", res.UserId)
		c.Next()
	}
}
