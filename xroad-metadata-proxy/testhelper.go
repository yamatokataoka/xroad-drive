package proxy

import (
  "testing"

  "github.com/alicebob/miniredis"
  "github.com/go-redis/redis"
)

func NewRedisMock(t *testing.T) (*redis.Client, *miniredis.Miniredis) {
  t.Helper()

  miniredis, err := miniredis.Run()
  if err != nil {
    t.Fatalf("Error while createing test redis server '%#v'", err)
  }

  client := redis.NewClient(&redis.Options{
    Addr:     miniredis.Addr(),
    Password: "",
    DB:       0,
  })

  return client, miniredis
}
