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

func verifyToken(tk string) (*jwt.Token, error) {
	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetConfig().Get("jwt.secret").(string)), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// will not check the token TTL since the TTL info was stored inside the redis
func ExtractTokenData(token string) (*JwtInfo, error) {
	tk, err := verifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := tk.Claims.(jwt.MapClaims)
	if ok && tk.Valid {
		account, ok1 := claims["account"].(string)
		if !ok1 {
			return nil, err
		}
		userId, ok2 := claims["userId"].(float64)
		if !ok2 {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		return &JwtInfo{
			Account: account,
			UserId:  int64(userId),
		}, nil
	}

	return nil, err
}
