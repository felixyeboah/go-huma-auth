package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type Store struct {
	redisClient *redis.Client
}

func NewStore(redisClient *redis.Client) *Store {
	return &Store{
		redisClient: redisClient,
	}
}

func (r *Store) StoreToken(ctx context.Context, userId, token string, duration time.Duration) error {
	err := r.redisClient.Set(ctx, "token:"+userId, token, duration).Err()
	if err != nil {
		return errors.New("failed to store token: " + err.Error())
	}

	return nil
}

func (r *Store) GetToken(ctx context.Context, userId string) (string, error) {
	token, err := r.redisClient.Get(ctx, "token:"+userId).Result()
	if err != nil {
		return "", errors.New("token not found")
	}

	return token, nil
}

func (r *Store) DeleteToken(ctx context.Context, userId string) error {
	err := r.redisClient.Del(ctx, "token:"+userId).Err()
	if err != nil {
		return errors.New("failed to delete token: " + err.Error())
	}

	return nil
}
