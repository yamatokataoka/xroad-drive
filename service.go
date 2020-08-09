package proxy

type ProviderService interface {
	GetAll() ([]*XRoadMember, error)
}

type providerService struct {
	pr ProviderRepository
}

func NewProviderService(pr ProviderRepository) ProviderService {
  return &providerService{pr}
}

func (ps *providerService) GetAll() ([]*XRoadMember, error) {
  providers, err := ps.pr.GetAll()
  if err != nil {
    return nil, err
  }
  return providers, nil
}