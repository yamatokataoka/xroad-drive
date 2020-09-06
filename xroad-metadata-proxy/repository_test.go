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
  // TODO: Replace error handling to assert.NoError(t, err)
  if err != nil {
    t.Errorf("Failed to call GetAll '%#v'", err)
  }

  assert.Equal(t, expectedXRoadMembers, actualXRoadMembers)
}

func TestProviderRepository_Set(t *testing.T) {
  mockClient, miniredis := NewRedisMock(t)
  defer miniredis.Close()

  xRoadMembers := decodeXRoadMembers(t, mapXRoadMembers)
  expectedXRoadMembers := xRoadMembers

  providerRepository := NewProviderRepository(mockClient)
  err := providerRepository.Set(xRoadMembers)
  assert.NoError(t, err)

  actualXRoadMembers, err := getXRoadMembersByMatch(mockClient, providerKey + ":*")
  if err != nil {
    t.Errorf("Failed to get actual result '%#v'", err)
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
  // TODO: Replace error handling to assert.NoError(t, err)
  if err != nil {
    t.Errorf("Failed to call GetAll '%#v'", err)
  }

  assert.Equal(t, expectedXRoadMembers, actualXRoadMembers)
}

func TestClientRepository_Set(t *testing.T) {
  mockClient, miniredis := NewRedisMock(t)
  defer miniredis.Close()

  xRoadMembers := decodeXRoadMembers(t, mapXRoadMembers)
  expectedXRoadMembers := xRoadMembers

  clientRepository := NewClientRepository(mockClient)
  err := clientRepository.Set(xRoadMembers)
  assert.NoError(t, err)

  actualXRoadMembers, err := getXRoadMembersByMatch(mockClient, clientKey + ":*")
  if err != nil {
    t.Errorf("Failed to get actual result '%#v'", err)
  }

  assert.Equal(t, expectedXRoadMembers, actualXRoadMembers)
}
