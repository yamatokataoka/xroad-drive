package proxy

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestProvider_GetAll(t *testing.T) {
  mockClient, miniredis := NewRedisMock(t)
  defer miniredis.Close()

  mapProviders := []map[string]interface{}{
    {"ID": "CS:ORG:1111:Company1Provider", "Name": "Company 1"},
    {"ID": "CS:ORG:1112:Company2Provider", "Name": "Company 2"},
    {"ID": "CS:ORG:1113:Company3Provider", "Name": "Company 3"},
  }

  expectedProviders := HMSetAll(t, mapProviders, mockClient, providerKey)

  providerRepository := NewProviderRepository(mockClient)
  actualProviders, err := providerRepository.GetAll()
  if err != nil {
    t.Errorf("Failed to call GetAll '%#v'", err)
  }

  assert.Equal(t, expectedProviders, actualProviders)
}