package proxy

import "errors"

type ProviderService interface {
	Get() ([]*XRoadMember, error)
}

type providerService struct {
	providerRepository ProviderRepository
}

func NewProviderService(pr ProviderRepository) ProviderService {
  return &providerService{pr}
}

func (ps *providerService) Get() ([]*XRoadMember, error) {
  return nil, errors.New("Not implemented")
}