package proxy

import (
  "errors"

  "github.com/go-redis/redis"
  log "github.com/sirupsen/logrus"
  "github.com/mitchellh/mapstructure"
)

var providerKey = "providers"
var clientKey = "clients"

type XRoadMemberRepository interface {
  GetAll() ([]*XRoadMember, error)
}

type ProviderRepository interface {
  XRoadMemberRepository
}

type ClientRepository interface {
  XRoadMemberRepository
}

type providerRepository struct {
	client *redis.Client
}

type clientRepository struct {
	client *redis.Client
}

func NewProviderRepository(c *redis.Client) ProviderRepository {
  return &providerRepository{c}
}

func NewClientRepository(c *redis.Client) ClientRepository {
  return &clientRepository{c}
}

func (pr *providerRepository) GetAll() ([]*XRoadMember, error) {
  match := providerKey + ":*"

  providers, err := getXRoadMembers(pr.client, match)
  if err != nil {
    return nil, errors.New("Failed to get X-Road member information")
  }

  return providers, nil
}

func (pr *clientRepository) GetAll() ([]*XRoadMember, error) {
  match := clientKey + ":*"

  clients, err := getXRoadMembers(pr.client, match)
  if err != nil {
    return nil, errors.New("Failed to get X-Road member information")
  }

  return clients, nil
}

func getXRoadMembers(client *redis.Client, match string) ([]*XRoadMember, error) {
  var (
    cursor uint64
    xRoadMembers []*XRoadMember
  )
  for {
    var keys []string
    var err error
    keys, cursor, err = client.Scan(cursor, match, 10).Result()
    if err != nil {
      log.
        WithField("match", match).
        Error("Failed to scan keys")
      return nil, err
    }
    for _, key := range keys {
      keyType, err := client.Type(key).Result()
      if err != nil {
        log.
          WithError(err).
          WithField("key", key).
          Error("Failed to check type")
        return nil, err
      }
      if keyType == "hash" {
        var xRoadMember XRoadMember

        mapXRoadMember, err := client.HGetAll(key).Result()
        if err != nil {
          log.
            WithError(err).
            WithField("key", key).
            Error("Failed to get all fields")
          return nil, err
        }

        err = mapstructure.Decode(mapXRoadMember, &xRoadMember)
        if err != nil {
          log.
            WithError(err).
            WithField("input", mapXRoadMember).
            WithField("output", xRoadMember).
            Error("Failed to decode")
          return nil, err
        }
        xRoadMembers = append(xRoadMembers, &xRoadMember)
      }
    }
    if cursor == 0 {
      break
    }
  }
  return xRoadMembers, nil
}