package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	// redisClient := NewRedisClient()

}

// 创建 redis 客户端
func NewRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis-service:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}

	return redisClient
}
