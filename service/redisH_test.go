package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedisHSetHGet(t *testing.T) {
	// key => clase, field => student name, value => student number
	k := "classA"
	f1, v1 := "jason", 13
	f2, v2 := "mike", 15

	redisHSet(testRedisClient, k, f1, v1)
	redisHSet(testRedisClient, k, f2, v2)

	value := redisHGet(testRedisClient, k, f1)
	require.Equal(t, strconv.Itoa(v1), value)

	value = redisHGet(testRedisClient, k, f2)
	require.Equal(t, strconv.Itoa(v2), value)
}

func TestRedisHGetAll(t *testing.T) {
	k := "classA"
	f1, v1 := "jason", 13
	f2, v2 := "mike", 15

	redisHSet(testRedisClient, k, f1, v1)
	redisHSet(testRedisClient, k, f2, v2)
	m := redisHGetAll(testRedisClient, k)
	for key, value := range m {
		if key == f1 {
			require.Equal(t, strconv.Itoa(v1), value)
		}
		if key == f2 {
			require.Equal(t, strconv.Itoa(v2), value)
		}
	}
}

func TestHDel(t *testing.T) {
	k := "classA"
	f1, v1 := "jason", 13

	redisHSet(testRedisClient, k, f1, v1)
	err := redisHDel(testRedisClient, k, f1)
	require.NoError(t, err)
	value := redisHGet(testRedisClient, k, f1)
	require.Equal(t, "", value)
}
