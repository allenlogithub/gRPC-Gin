package databases

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type (
	UserAccessToken struct {
		UserId      int64
		AccessToken string
		TTL         int
	}

	UserOwnedToken struct {
		AccessToken string
	}
)

var (
	ctx = context.Background()
)

// key = accessToken, val = userAccount
func SetAuthToken(d *UserAccessToken) error {
	if err := connRedis.Set(ctx, d.AccessToken, d.UserId, time.Duration(d.TTL)*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

// pair with SetAuthToken
// returns redis.Nil error when key does not exist
func GetUserAccount(d *UserOwnedToken) (string, error) {
	res := connRedis.Get(ctx, d.AccessToken).Val()
	if err == redis.Nil {
		return "", errors.New("RedisGetFailed")
	}

	return res, nil
}

func DeleteAuthToken(d *UserOwnedToken) error {
	if err := connRedis.Del(ctx, d.AccessToken).Err(); err != nil {
		return err
	}

	return nil
}
