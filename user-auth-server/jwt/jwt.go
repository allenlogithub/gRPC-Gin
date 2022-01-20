package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	config "user-auth-server/config"
)

type (
	JwtInfo struct {
		Account string
		UserId  int64
		rand    string
	}
)

func CreateToken(j *JwtInfo) (string, error) {
	c := config.GetConfig()
	hmacSecret := []byte(c.Get("jwt.secret").(string))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"account": j.Account,
		"userId":  j.UserId,
		"rand":    uuid.New().String(),
	})
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
