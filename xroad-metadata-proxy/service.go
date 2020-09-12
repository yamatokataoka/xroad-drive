package proxy

import "strings"

type ServiceProviderService interface {
  FindServiceProviders(serviceClientID string, serviceCode string) ([]*XRoadMember, error)
}

type ServiceClientService interface {
  GetServiceClientsByService(serviceID string) ([]*XRoadMember, error)
}

type serviceProviderService struct {
  spr ServiceProviderRepository
  sr  ServiceRepository
}

type serviceClientService struct {
  scr ServiceClientRepository
}

func NewServiceProviderService(spr ServiceProviderRepository, sr  ServiceRepository) ServiceProviderService {
  return &serviceProviderService{spr, sr}
}

func NewServiceClientService(scr ServiceClientRepository) ServiceClientService {
  return &serviceClientService{scr}
}

func (sps *serviceProviderService) FindServiceProviders(serviceClientID string, serviceCode string) ([]*XRoadMember, error) {
  allProviders, err := sps.spr.GetAllServiceProviders()
  if err != nil {
    return nil, err
  }

  if len(serviceCode) == 0 {
    return allProviders, nil
  }

  allowedProviders := make([]*XRoadMember, 0)

  for _, provider := range allProviders {
    allowedServices, err := sps.sr.GetAllowedServices(serviceClientID, provider)
    if err != nil {
      continue
    }

    for _, allowedService := range allowedServices {
      allowedServiceCode := allowedService.ID[strings.LastIndex(allowedService.ID, ":")+1:]

      if allowedServiceCode == serviceCode {
        allowedProviders = append(allowedProviders, provider)
      }
    }
  }

  return allowedProviders, nil
}

func (scs *serviceClientService) GetServiceClientsByService(serviceID string) ([]*XRoadMember, error) {
  serviceClients, err := scs.scr.GetServiceClientsByService(serviceID)
  if err != nil {
    return nil, err
  }

  return serviceClients, nil
}
