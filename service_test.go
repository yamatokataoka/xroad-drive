package proxy

import (
  "testing"

  "github.com/stretchr/testify/mock"
  "github.com/stretchr/testify/assert"
)

type MockProviderRepository struct {
  mock.Mock
}

type MockClientRepository struct {
  mock.Mock
}

func (m *MockProviderRepository) GetAll() ([]*XRoadMember, error) {
  args := m.Called()
  return args.Get(0).([]*XRoadMember), args.Error(1)
}

func (m *MockClientRepository) GetAll() ([]*XRoadMember, error) {
  args := m.Called()
  return args.Get(0).([]*XRoadMember), args.Error(1)
}

func TestService_Provider_GetAll(t *testing.T) {
  expectedProviders := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Provider", Name: "Company 1"},
    &XRoadMember{ID: "CS:ORG:1112:Company2Provider", Name: "Company 2"},
  }

  mockProviderRepo := new(MockProviderRepository)

  mockProviderRepo.On("GetAll").Return(expectedProviders, nil)

  providerService := NewProviderService(mockProviderRepo)
  actualProviders, err := providerService.GetAll()
  if err != nil {
    t.Errorf("Failed to call GetAll '%#v'", err)
  }

  assert.Equal(t, expectedProviders, actualProviders)
}

func TestService_Client_GetAll(t *testing.T) {
  expectedClients := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Provider", Name: "Company 1"},
    &XRoadMember{ID: "CS:ORG:1112:Company2Provider", Name: "Company 2"},
  }

  mockClientRepo := new(MockClientRepository)

  mockClientRepo.On("GetAll").Return(expectedClients, nil)

  clientService := NewClientService(mockClientRepo)
  actualClients, err := clientService.GetAll()
  if err != nil {
    t.Errorf("Failed to call GetAll '%#v'", err)
  }

  assert.Equal(t, expectedClients, actualClients)
}