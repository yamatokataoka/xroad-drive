package proxy

import (
  "github.com/go-redis/redis"
  log "github.com/sirupsen/logrus"
  "github.com/mitchellh/mapstructure"
)

var providerKey = "providers"

type ProviderRepository interface {
  GetAll() ([]*XRoadMember, error)
}

type providerRepository struct {
	client *redis.Client
}

func NewProviderRepository(c *redis.Client) ProviderRepository {
  return &providerRepository{c}
}

func (pr *providerRepository) GetAll() ([]*XRoadMember, error) {
  var (
    providers []*XRoadMember
    allKeys []string
    match = providerKey + ":*"
    cursor uint64
  )

  for {
    var keys []string
    var err error
    keys, cursor, err = pr.client.Scan(cursor, match, 10).Result()
    if err != nil {
      log.
        WithField("match", match).
        Error("Failed to scan keys")
      return nil, err
    }
    allKeys = append(allKeys, keys...)
    if cursor == 0 {
      break
    }
  }

  for _, key := range allKeys {
    var xRoadMember XRoadMember
    mapProvider, err := pr.client.HGetAll(key).Result()
    if err != nil {
      log.
        WithError(err).
        WithField("key", key).
        Error("Failed to get all fields")
      continue
    }
    err = mapstructure.Decode(mapProvider, &xRoadMember)
    if err != nil {
      log.
        WithError(err).
        WithField("input", mapProvider).
        WithField("output", xRoadMember).
        Error("Failed to decode")
      continue
    }
    providers = append(providers, &xRoadMember)
  }

  return providers, nil
}