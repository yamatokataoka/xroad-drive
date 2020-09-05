package proxy

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

var mapXRoadMembers = []map[string]interface{}{
  {"ID": "CS:ORG:1111:Company1Subsystem", "Name": "Company 1"},
  {"ID": "CS:ORG:1112:Company2Subsystem", "Name": "Company 2"},
  {"ID": "CS:ORG:1113:Company3Subsystem", "Name": "Company 3"},
}

func TestProviderRepository_GetAll(t *testing.T) {
  mockClient, miniredis := NewRedisMock(t)
  defer miniredis.Close()

  HMSetAll(t, mapXRoadMembers, mockClient, providerKey)
  expectedXRoadMembers := decodeXRoadMembers(t, mapXRoadMembers)

  providerRepository := NewProviderRepository(mockClient)
  actualXRoadMembers, err := providerRepository.GetAll()
  if err != nil {
    t.Errorf("Failed to call GetAll '%#v'", err)
  }

  assert.Equal(t, expectedXRoadMembers, actualXRoadMembers)
}

func TestClientRepository_GetAll(t *testing.T) {
  mockClient, miniredis := NewRedisMock(t)
  defer miniredis.Close()

  HMSetAll(t, mapXRoadMembers, mockClient, clientKey)
  expectedXRoadMembers := decodeXRoadMembers(t, mapXRoadMembers)

  clientRepository := NewClientRepository(mockClient)
  actualXRoadMembers, err := clientRepository.GetAll()
  if err != nil {
    t.Errorf("Failed to call GetAll '%#v'", err)
  }

  assert.Equal(t, expectedXRoadMembers, actualXRoadMembers)
}
