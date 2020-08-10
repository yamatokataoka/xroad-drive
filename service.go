package proxy

import "errors"

type XRoadMemberService interface {
  GetAll() ([]*XRoadMember, error)
}

type ProviderService interface {
	XRoadMemberService
}

type ClientService interface {
	XRoadMemberService
}

type providerService struct {
	pr ProviderRepository
}

type clientService struct {
	cr ClientRepository
}

func NewProviderService(pr ProviderRepository) ProviderService {
  return &providerService{pr}
}

func NewClientService(cr ClientRepository) ClientService {
  return &clientService{cr}
}

func (ps *providerService) GetAll() ([]*XRoadMember, error) {
  providers, err := ps.pr.GetAll()
  if err != nil {
    return nil, err
  }
  return providers, nil
}

func (cs *clientService) GetAll() ([]*XRoadMember, error) {
  return nil, errors.New("Not implemented")
}