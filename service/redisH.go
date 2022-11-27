package main

import (
	"github.com/go-redis/redis"
)

func redisHSet(client *redis.Client, key, field string, value interface{}) {
	err := client.HSet(key, field, value).Err()
	if err != nil {
		panic(err)
	}
}

func redisHGetAll(client *redis.Client, key string) map[string]string {

	// 一次性返回 key 的所有 hash 字段和值
	// 回傳是一個 map
	return client.HGetAll(key).Val()
}

func redisHGet(client *redis.Client, key, field string) string {
	// key 是 hash key，可以想像成是 key value 中有 key value
	return client.HGet(key, field).Val()
}

func redisHDel(client *redis.Client, key string, field string) error {
	// HDel 可以支援一次放入多個 field ，不過此範例只示範一個 field
	return client.HDel(key, field).Err()
}
