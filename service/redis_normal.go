package main

import (
	"time"

	"github.com/go-redis/redis"
)

func redisSet(client *redis.Client, key, value string, expiration time.Duration) error {
	return client.Set(key, value, expiration).Err()
}

func redisGet(client *redis.Client, key string) string {
	return client.Get(key).Val()
}

// 如果這個 key 已經存在，則不會再設定新的上去
func redisGetSet(client *redis.Client, key, value string) (string, error) {
	return client.GetSet(key, value).Result()
}

// SetNX => 如果這個 key 不存在，就設置這個 key, value
// MGet => 一次查詢多個 key 的 value
// MSet => 一次設置多個 key, value
// Expire => 設定 key 的過期時間
