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
  Set(xRoadMembers []*XRoadMember) error
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

  providers, err := getXRoadMembersByMatch(pr.client, match)
  if err != nil {
    return nil, errors.New("Failed to get X-Road member information")
  }

  return providers, nil
}

func (pr *providerRepository) Set(xRoadMembers []*XRoadMember) error {
  mapXRoadMembers, err := decodeStructs(xRoadMembers)
  if err != nil {
    return err
  }

  err = hMSetXRoadMembers(pr.client, mapXRoadMembers, providerKey)
  if err != nil {
    return err
  }

  return nil
}

func (cr *clientRepository) GetAll() ([]*XRoadMember, error) {
  match := clientKey + ":*"

  clients, err := getXRoadMembersByMatch(cr.client, match)
  if err != nil {
    return nil, errors.New("Failed to get X-Road member information")
  }

  return clients, nil
}

func (cr *clientRepository) Set(xRoadMembers []*XRoadMember) error {
  mapXRoadMembers, err := decodeStructs(xRoadMembers)
  if err != nil {
    return err
  }

  err = hMSetXRoadMembers(cr.client, mapXRoadMembers, clientKey)
  if err != nil {
    return err
  }

  return nil
}

func getXRoadMembersByMatch(client *redis.Client, match string) ([]*XRoadMember, error) {
  var (
    cursor uint64
    xRoadMembers = make([]*XRoadMember, 0)
  )

  for {
    var keys []string
    var err error
    keys, cursor, err = client.Scan(cursor, match, 10).Result()
    if err != nil {
      log.
        WithError(err).
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
            WithFields(log.Fields{
              "input":  mapXRoadMember,
              "output": xRoadMember,
            }).
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

func hMSetXRoadMembers(client *redis.Client, mapXRoadMembers []map[string]interface{}, key string) error {
  for _, mapXRoadMember := range mapXRoadMembers {
    key = key + ":" + mapXRoadMember["ID"].(string)

    err := client.HMSet(key, mapXRoadMember).Err()
    if err != nil {
      log.
        WithError(err).
        WithFields(log.Fields{
          "key":  key,
          "values": mapXRoadMember,
        }).
        Error("Failed to set data")
      return err
    }
  }

  return nil
}

func decodeStructs(xRoadMembers []*XRoadMember) ([]map[string]interface{}, error) {
  var mapXRoadMembers []map[string]interface{}

  for _, xRoadMember := range xRoadMembers {
    var mapXRoadMember map[string]interface{}

    err := mapstructure.Decode(xRoadMember, &mapXRoadMember)
    if err != nil {
      log.
        WithError(err).
        WithFields(log.Fields{
          "input":  xRoadMember,
          "output": mapXRoadMember,
        }).
        Error("Failed to decode")
      return nil, err
    }

    mapXRoadMembers = append(mapXRoadMembers, mapXRoadMember)
  }

  return mapXRoadMembers, nil
}
