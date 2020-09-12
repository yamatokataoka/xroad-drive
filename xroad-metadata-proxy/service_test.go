package proxy

import (
  "testing"

  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/assert"
)

type MockServiceProviderRepository struct {
  mock.Mock
}

type MockServiceRepository struct {
  mock.Mock
}

type MockServiceClientRepository struct {
  mock.Mock
}

func (m *MockServiceProviderRepository) GetAllServiceProviders() ([]*XRoadMember, error) {
  args := m.Called()
  return args.Get(0).([]*XRoadMember), args.Error(1)
}

func (m *MockServiceRepository) GetAllowedServices(serviceClientID string, serviceProvider *XRoadMember) ([]*Service, error) {
  args := m.Called(serviceClientID, serviceProvider)
  return args.Get(0).([]*Service), args.Error(1)
}

func (m *MockServiceClientRepository) GetServiceClientsByService(serviceID string) ([]*XRoadMember, error) {
  args := m.Called(serviceID)
  return args.Get(0).([]*XRoadMember), args.Error(1)
}

func TestServiceProviderService_FindServiceProviders(t *testing.T) {
  allProviders := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Subsystem", Name: "Company 1"},
    &XRoadMember{ID: "CS:ORG:1112:Company2Subsystem", Name: "Company 2"},
  }

  expectedProviders := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Subsystem", Name: "Company 1"},
  }

  testServices := []*Service{
    &Service{ID: "CS:ORG:1111:Company1Subsystem:TestService"},
  }
  otherServices := []*Service{
    &Service{ID: "CS:ORG:1112:Company2Subsystem:OtherService"},
  }

  mockspr := new(MockServiceProviderRepository)
  mocksr := new(MockServiceRepository)

  mockspr.
    On("GetAllServiceProviders").
    Return(allProviders, nil)

  provider := &XRoadMember{ID: "CS:ORG:1111:Company1Subsystem", Name: "Company 1"}
  mocksr.
    On("GetAllowedServices", mock.Anything, provider).
    Return(testServices, nil)
  mocksr.
    On("GetAllowedServices", mock.Anything, mock.Anything).
    Return(otherServices, nil)

  serviceClientID := "Anything:Anything:Anything:Anything"
  serviceCode := "TestService"

  sps := NewServiceProviderService(mockspr, mocksr)
  actualProviders, err := sps.FindServiceProviders(serviceClientID, serviceCode)
  assert.NoError(t, err)

  assert.Equal(t, expectedProviders, actualProviders)
}

func TestServiceClientService_GetServiceClientsByService(t *testing.T) {
  expectedClients := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Provider", Name: "Company 1"},
    &XRoadMember{ID: "CS:ORG:1112:Company2Provider", Name: "Company 2"},
  }

  mockscr := new(MockServiceClientRepository)

  serviceID := "CS:ORG:1111:Company1Provider:TestService"

  mockscr.
    On("GetServiceClientsByService", serviceID).
    Return(expectedClients, nil)

  scs := NewServiceClientService(mockscr)
  actualClients, err := scs.GetServiceClientsByService(serviceID)
  assert.NoError(t, err)

  assert.Equal(t, expectedClients, actualClients)
}
