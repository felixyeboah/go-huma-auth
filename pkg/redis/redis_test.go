package redis

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type User struct {
	userId   string        `json:"userId"`
	token    string        `json:"token"`
	duration time.Duration `json:"duration"`
}

var args = User{
	userId:   "user1",
	token:    uuid.New().String(),
	duration: time.Hour,
}

func TestStoreToken(t *testing.T) {
	client := NewRedisClient()
	redisClient := NewStore(client)

	err := redisClient.StoreToken(ctx, args.userId, args.token, args.duration)
	assert.NoError(t, err)
}

func TestGetToken(t *testing.T) {
	client := NewRedisClient()
	redisClient := NewStore(client)

	token, err := redisClient.GetToken(ctx, args.userId)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGetTokenExpired(t *testing.T) {
	client := NewRedisClient()
	redisClient := NewStore(client)

	token, err := redisClient.GetToken(ctx, uuid.New().String())
	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestDeleteToken(t *testing.T) {
	client := NewRedisClient()
	redisClient := NewStore(client)

	err := redisClient.DeleteToken(ctx, args.userId)
	assert.NoError(t, err)
}
