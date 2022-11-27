package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	k := "cafeA"
	v := "good"
	e := 5 * time.Second
	err := redisSet(testRedisClient, k, v, e)
	require.NoError(t, err)

	value := redisGet(testRedisClient, k)
	require.Equal(t, v, value)
}

func TestRedisGetSet(t *testing.T) {
	k := "cafeBB"
	v := "so so"

	s, err := redisGetSet(testRedisClient, k, v)
	require.NoError(t, err)
	require.Equal(t, v, s)

	newV := "good"
	err = redisSet(testRedisClient, k, newV, time.Second*10)
	require.NoError(t, err)

	s, err = redisGetSet(testRedisClient, k, v) // 這邊還是設定 v
	require.NoError(t, err)
	require.Equal(t, newV, s) // 如果 key 已經存在，則不會再次設定
}
