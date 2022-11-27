package main

import (
	"os"
	"testing"

	"github.com/go-redis/redis"
)

var testRedisClient *redis.Client

func TestMain(m *testing.M) {
	testRedisClient = NewRedisClient()
	os.Exit(m.Run())
}
