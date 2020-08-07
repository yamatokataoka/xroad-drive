package proxy

import (
  "testing"
  "os"

  "github.com/stretchr/testify/assert"
  "github.com/mitchellh/mapstructure"
  log "github.com/sirupsen/logrus"
)

var mapProviders = []map[string]interface{}{
  {"ID": "CS:ORG:1111:Company1Provider", "Name": "Company 1"},
  {"ID": "CS:ORG:1112:Company2Provider", "Name": "Company 2"},
  {"ID": "CS:ORG:1113:Company3Provider", "Name": "Company 3"},
}

var expectedProviders []*XRoadMember

func TestMain(m *testing.M) {
  for _, mapProvider := range mapProviders {
    var xRoadMember XRoadMember
    err := mapstructure.Decode(mapProvider, &xRoadMember)
    if err != nil {
      log.Fatal(err)
    }
    expectedProviders = append(expectedProviders, &xRoadMember)
  }

	code := m.Run()
	os.Exit(code)
}

func TestProvider_Get(t *testing.T) {
  mockClient, miniredis := NewRedisMock(t)
  defer miniredis.Close()

  for _, mapProvider := range mapProviders {
    err := mockClient.HMSet("providers:" + mapProvider["ID"].(string), mapProvider).Err()
    if err != nil {
      t.Fatalf("Error while setting test data '%#v'", err)
    }
  }

  providerRepository := NewProviderRepository(mockClient)
  actualProviders, err := providerRepository.GetAll()
  if err != nil {
    t.Errorf("Error while calling GetAll function '%#v'", err)
  }

  assert.Equal(t, expectedProviders, actualProviders)
}