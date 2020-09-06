package proxy

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

var expectedMapXRoadMembers = []map[string]interface{}{
  {"ID": "CS:ORG:1111:Company1Subsystem", "Name": "Company 1"},
  {"ID": "CS:ORG:1112:Company2Subsystem", "Name": "Company 2"},
  {"ID": "CS:ORG:1113:Company3Subsystem", "Name": "Company 3"},
}

var expectedXRoadMembers = []*XRoadMember{
  {ID: "CS:ORG:1111:Company1Subsystem", Name: "Company 1"},
  {ID: "CS:ORG:1112:Company2Subsystem", Name: "Company 2"},
  {ID: "CS:ORG:1113:Company3Subsystem", Name: "Company 3"},
}

func TestProviderRepository_GetAll(t *testing.T) {
  mockClient, miniredis := NewRedisMock(t)
  defer miniredis.Close()

  for _, expectedMapXRoadMember := range expectedMapXRoadMembers {
    err := hMSetXRoadMember(mockClient, expectedMapXRoadMember, providerKey)
    if err != nil {
      t.Errorf("Failed to set test data '%#v'", err)
    }
  }

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

  providerRepository := NewProviderRepository(mockClient)
  err := providerRepository.Set(expectedXRoadMembers)
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

  for _, expectedMapXRoadMember := range expectedMapXRoadMembers {
    err := hMSetXRoadMember(mockClient, expectedMapXRoadMember, clientKey)
    if err != nil {
      t.Errorf("Failed to set test data '%#v'", err)
    }
  }

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

  clientRepository := NewClientRepository(mockClient)
  err := clientRepository.Set(expectedXRoadMembers)
  assert.NoError(t, err)

  actualXRoadMembers, err := getXRoadMembersByMatch(mockClient, clientKey + ":*")
  if err != nil {
    t.Errorf("Failed to get actual result '%#v'", err)
  }

  assert.Equal(t, expectedXRoadMembers, actualXRoadMembers)
}
