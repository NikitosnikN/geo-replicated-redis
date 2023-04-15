package redis

import "github.com/redis/go-redis/v9"

func NewClient(connectionString string) (*redis.Client, error) {
	uri, err := redis.ParseURL(connectionString)

	if err != nil {
		return nil, err
	}

	return redis.NewClient(uri), nil
}
