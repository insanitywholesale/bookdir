package redis

import (
	"github.com/go-redis/redis"
	pb "gitlab.com/insanitywholesale/proto/v1"
)

type redisRepo struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepo(redisURL string) (*redisRepo, error) {
	redisclient, err := newRedisClient(redisURL)
	if err != nil {
		return nil, err
	}
	return &repo{client: redisclient}, nil
}

