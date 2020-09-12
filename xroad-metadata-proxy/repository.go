package proxy

import (
  "errors"
  "strings"
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"

  log "github.com/sirupsen/logrus"
)

type ServiceProviderRepository interface {
  GetAllServiceProviders() ([]*XRoadMember, error)
}

type ServiceRepository interface {
  GetAllowedServices(serviceClientID string, serviceProvider *XRoadMember) ([]*Service, error)
}

type ServiceClientRepository interface {
  GetServiceClientsByService(serviceID string) ([]*XRoadMember, error)
}

type serviceProviderRepository struct {
  c *MetadataServiceClient
}

type serviceRepository struct {
  c *MetadataServiceClient
}

type serviceClientRepository struct {
  c *AdminAPIClient
}

func NewServiceProviderRepository(c *MetadataServiceClient) ServiceProviderRepository {
  return &serviceProviderRepository{c}
}

func NewServiceRepository(c *MetadataServiceClient) ServiceRepository {
  return &serviceRepository{c}
}

func NewServiceClientRepository(c *AdminAPIClient) ServiceClientRepository {
  return &serviceClientRepository{c}
}

func (spr *serviceProviderRepository) GetAllServiceProviders() ([]*XRoadMember, error) {
  url := fmt.Sprintf("%s/listClients", spr.c.BaseURL)

  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to create http request")
    return nil, err
  }

  req.Header.Set("Accept", "application/json")

  res, err := spr.c.HTTPClient.Do(req)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to send")
    return nil, err
  }

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to read")
    return nil, err
  }

  if res.StatusCode != 200 {
    log.
      WithFields(log.Fields{
        "status_code": res.StatusCode,
        "body":        string(body),
      }).
      Error("Failed to retrieve service providers from Metadata Service")

    return nil, errors.New("Failed to retrieve service providers from Metadata Service")
  }

  var unmarshaledBody map[string]interface{}
  err = json.Unmarshal(body, &unmarshaledBody)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to unmarshal")
    return nil, err
  }

  allProviders := make([]*XRoadMember, 0)
  for _, member := range unmarshaledBody["member"].([]interface{}) {
    id := member.(map[string]interface{})["id"].(map[string]interface{})

    _, found := id["subsystem_code"].(string)
    if !found {
      continue
    }

    providerID := fmt.Sprintf(
      "%s:%s:%s:%s",
      id["xroad_instance"].(string),
      id["member_class"].(string),
      id["member_code"].(string),
      id["subsystem_code"].(string),
    )

    xRoadMember := XRoadMember{}
    xRoadMember.ID = providerID
    xRoadMember.Name = member.(map[string]interface{})["name"].(string)

    allProviders = append(allProviders, &xRoadMember)
  }

  return allProviders, nil
}

func (sr *serviceRepository) GetAllowedServices(serviceClientID string, serviceProvider *XRoadMember) ([]*Service, error) {
  providerPathID := strings.ReplaceAll(serviceProvider.ID, ":", "/")
  url := fmt.Sprintf("%s/r1/%s/allowedMethods", sr.c.BaseURL, providerPathID)

  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to create http request")
    return nil, err
  }

  clientPathID := strings.ReplaceAll(serviceClientID, ":", "/")
  req.Header.Set("X-Road-Client", clientPathID)

  res, err := sr.c.HTTPClient.Do(req)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to send")
    return nil, err
  }

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to read")
    return nil, err
  }

  if res.StatusCode != 200 {
    log.
      WithFields(log.Fields{
        "status_code": res.StatusCode,
        "body":   string(body),
      }).
      Error("Failed to retrieve allowed services from Metadata Service")

    return nil, errors.New("Failed to retrieve allowed services from Metadata Service")
  }

  var unmarshaledBody map[string]interface{}
  err = json.Unmarshal(body, &unmarshaledBody)
  if err != nil {
    log.
    WithError(err).
    Error("Failed to unmarshal")
    return nil, err
  }

  allowedServices := make([]*Service, 0)
  for _, service := range unmarshaledBody["service"].([]interface{}) {
    serviceID := fmt.Sprintf(
      "%s:%s:%s:%s:%s",
      service.(map[string]interface{})["xroad_instance"],
      service.(map[string]interface{})["member_class"],
      service.(map[string]interface{})["member_code"],
      service.(map[string]interface{})["subsystem_code"],
      service.(map[string]interface{})["service_code"],
    )

    service := Service{ID: serviceID}

    allowedServices = append(allowedServices, &service)
  }

  return allowedServices, nil
}

func (scr *serviceClientRepository) GetServiceClientsByService(serviceID string) ([]*XRoadMember, error) {
  url := fmt.Sprintf("%s/services/%s/service-clients", scr.c.BaseURL, serviceID)

  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to create http request")
    return nil, err
  }

  apiKey := fmt.Sprintf("X-Road-ApiKey token=%s", scr.c.apiKey)
  req.Header.Set("Authorization", apiKey)

  res, err := scr.c.HTTPClient.Do(req)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to send")
    return nil, err
  }

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to read")
    return nil, err
  }

  if res.StatusCode != 200 {
    log.
      WithFields(log.Fields{
        "status_code": res.StatusCode,
        "body":        string(body),
      }).
      Error("Failed to retrieve service clients from Admin API")

    return nil, errors.New("Failed to retrieve service clients from Admin API")
  }

  var clients []*XRoadMember
  err = json.Unmarshal(body, &clients)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to unmarshal")
    return nil, err
  }

  return clients, nil
}
