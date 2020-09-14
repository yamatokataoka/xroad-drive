package proxy

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "encoding/json"
  "fmt"

  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/assert"
  "github.com/gorilla/mux"
)

type mockServiceProviderService struct {
  mock.Mock
}

type mockServiceClientService struct {
  mock.Mock
}

func (m *mockServiceProviderService) FindServiceProviders(serviceClientID string, serviceCode string) ([]*XRoadMember, error) {
  args := m.Called(serviceClientID, serviceCode)
  return args.Get(0).([]*XRoadMember), args.Error(1)
}

func (m *mockServiceClientService) GetServiceClientsByService(serviceID string) ([]*XRoadMember, error) {
  args := m.Called(serviceID)
  return args.Get(0).([]*XRoadMember), args.Error(1)
}

func TestGetServiceProviders(t *testing.T) {
  var actualProviders []*XRoadMember
  // TODO: Move it to test helpers
  expectedProviders := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Provider", Name: "Company 1"},
    &XRoadMember{ID: "CS:ORG:1112:Company2Provider", Name: "Company 2"},
  }

  serviceCode := "TestService"
  clientPathID := "CS/ORG/1111/Company1Subsystem"
  clientID := "CS:ORG:1111:Company1Subsystem"

  mocksps := new(mockServiceProviderService)
  mocksps.
    On("FindServiceProviders", clientID, serviceCode).
    Return(expectedProviders, nil)

  sph := NewServiceProviderHandler(mocksps)

  path := fmt.Sprintf("/service-providers?service-code=%s", serviceCode)
  req := httptest.NewRequest(http.MethodGet, path, nil)

  req.Header.Set("X-Road-Client", clientPathID)

  w := httptest.NewRecorder()

  http.HandlerFunc(sph.GetServiceProviders).ServeHTTP(w, req)

  json.NewDecoder(w.Body).Decode(&actualProviders)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Equal(t, expectedProviders, actualProviders)
  // TODO: Add more asserts
}

func TestGetServiceServiceClients(t *testing.T) {
  var actualClients []*XRoadMember
  // TODO: Move it to test helpers
  expectedClients := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Provider", Name: "Company 1"},
    &XRoadMember{ID: "CS:ORG:1112:Company2Provider", Name: "Company 2"},
  }

  serviceID := "CS:ORG:1111:Company1Subsystem:TestService"

  mockscs := new(mockServiceClientService)
  mockscs.
    On("GetServiceClientsByService", serviceID).
    Return(expectedClients, nil)

  sh := NewServiceHandler(mockscs)

  path := fmt.Sprintf("/services/%s/service-clients", serviceID)
  req := httptest.NewRequest(http.MethodGet, path, nil)

  w := httptest.NewRecorder()

  router := mux.NewRouter()
  router.HandleFunc("/services/{service-id}/service-clients", sh.GetServiceServiceClients)
  router.ServeHTTP(w, req)

  json.NewDecoder(w.Body).Decode(&actualClients)

  assert.Equal(t, http.StatusOK, w.Code)
  assert.Equal(t, expectedClients, actualClients)
  // TODO: Add more asserts
}
