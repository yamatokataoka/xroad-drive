package proxy

import (
  "time"

  "github.com/go-redis/redis"
  log "github.com/sirupsen/logrus"
)

func NewClient() *redis.Client {
  client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
  })

  // TODO: handle and log in the case, connected redis -> disconnected
  go retry(10*time.Second, func() (err error) {
    _, err = client.Ping().Result()
    if err != nil {
      log.
      WithError(err).
      Error("Failed to connect redis")
    } else {
      log.Info("Connected to redis")
    }
    return err
  })

  return client
}

func retry(sleep time.Duration, f func() error) {
  for {
    err := f()
    if err == nil {
      return
    }

    time.Sleep(sleep)
  }
}