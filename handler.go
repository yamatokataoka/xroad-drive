package proxy

import (
  "net/http"

  log "github.com/sirupsen/logrus"
)

type ProviderHandler interface {
  GetServiceProviders(w http.ResponseWriter, r *http.Request)
}

type providerHandler struct {
  providerService ProviderService
}

func NewProviderHandler(ps ProviderService) ProviderHandler {
  return &providerHandler{ps}
}

func (ph *providerHandler) GetServiceProviders(w http.ResponseWriter, r *http.Request) {
  log.Error("Not implemented")
}