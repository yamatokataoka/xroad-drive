package proxy

import (
  "errors"

  "github.com/go-redis/redis"
)

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
  return nil, errors.New("Not implemented")
}