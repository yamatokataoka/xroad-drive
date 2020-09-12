package proxy

import (
  "testing"
  "io/ioutil"
  "bytes"
  "net/http"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
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

type MockHTTPClient struct {
  mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
  args := m.Called(req)
  return args.Get(0).(*http.Response), args.Error(1)
}

func TestServiceProviderRepository_GetAllServiceProviders(t *testing.T) {
  mockHTTPClient := new(MockHTTPClient)

  responseJson := `{
  "member": [
      {
        "id": {
          "member_class": "ORG",
          "member_code": "1111",
          "object_type": "MEMBER",
          "xroad_instance": "CS"
        },
        "name": "Company 1"
      },
      {
        "id": {
          "member_class": "ORG",
          "member_code": "1111",
          "subsystem_code": "Company1Subsystem",
          "object_type": "SUBSYSTEM",
          "xroad_instance": "CS"
        },
        "name": "Company 1"
      }
    ]
  }`

  r := ioutil.NopCloser(bytes.NewReader([]byte(responseJson)))
  response := &http.Response{
    StatusCode: 200,
    Body:       r,
  }

  mockHTTPClient.
    On("Do", mock.Anything).
    Return(response, nil)

  c := &MetadataServiceClient{HTTPClient: mockHTTPClient}

  spr := NewServiceProviderRepository(c)

  expectedProviders := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Subsystem", Name: "Company 1"},
  }

  actualProviders, err := spr.GetAllServiceProviders()
  assert.NoError(t, err)

  assert.Equal(t, expectedProviders, actualProviders)
}

func TestServiceRepository_GetAllowedServices(t *testing.T) {
  mockHTTPClient := new(MockHTTPClient)

  responseJson := `{
    "service": [
      {
        "member_class": "ORG",
        "member_code": "1111",
        "object_type": "SERVICE",
        "service_code": "TestService",
        "subsystem_code": "Company1Subsystem",
        "xroad_instance": "CS"
      }
    ]
  }`

  r := ioutil.NopCloser(bytes.NewReader([]byte(responseJson)))
  response := &http.Response{
    StatusCode: 200,
    Body:       r,
  }

  mockHTTPClient.
    On("Do", mock.Anything).
    Return(response, nil)

  c := &MetadataServiceClient{HTTPClient: mockHTTPClient}

  sr := NewServiceRepository(c)

  expectedServices := []*Service{
    &Service{ID: "CS:ORG:1111:Company1Subsystem:TestService"},
  }

  actualServices, err := sr.GetAllowedServices("serviceClientID", new(XRoadMember))
  assert.NoError(t, err)

  assert.Equal(t, expectedServices, actualServices)
}

func TestServiceClientRepository_GetServiceClientsByService(t *testing.T) {
  mockHTTPClient := new(MockHTTPClient)

  responseJson := `[
    {
      "id": "CS:ORG:1111:Company1Subsystem",
      "name": "Company 1",
      "service_client_type": "SUBSYSTEM",
      "rights_given_at": "2018-12-15T00:00:00.001Z"
    }
  ]`

  r := ioutil.NopCloser(bytes.NewReader([]byte(responseJson)))
  response := &http.Response{
    StatusCode: 200,
    Body:       r,
  }

  mockHTTPClient.
    On("Do", mock.Anything).
    Return(response, nil)

  c := &AdminAPIClient{HTTPClient: mockHTTPClient}

  scs := NewServiceClientRepository(c)

  expectedClients := []*XRoadMember{
    &XRoadMember{ID: "CS:ORG:1111:Company1Subsystem", Name: "Company 1"},
  }

  actualClients, err := scs.GetServiceClientsByService("serviceID")
  assert.NoError(t, err)

  assert.Equal(t, expectedClients, actualClients)
}
