package proxy

import (
  "net/http"
  "encoding/json"
  "strings"

  log "github.com/sirupsen/logrus"
  "github.com/gorilla/mux"
)

type ServiceProviderHandler interface {
  GetServiceProviders(w http.ResponseWriter, r *http.Request)
}

type ServiceHandler interface {
  GetServiceServiceClients(w http.ResponseWriter, r *http.Request)
}

type serviceProviderHandler struct {
  sps ServiceProviderService
}

type serviceHandler struct {
  scs ServiceClientService
}

func NewServiceProviderHandler(sps ServiceProviderService) ServiceProviderHandler {
  return &serviceProviderHandler{sps}
}

func NewServiceHandler(scs ServiceClientService) ServiceHandler {
  return &serviceHandler{scs}
}

func (sph *serviceProviderHandler) GetServiceProviders(w http.ResponseWriter, r *http.Request) {
  log.
    WithFields(log.Fields{
      "method": r.Method,
      "path":   r.URL,
    }).
    Info("Receive request")

  queries := r.URL.Query()
  serviceCode := queries.Get("service-code")

  clientPathID := r.Header.Get("X-Road-Client")

  if len(clientPathID) == 0 {
    log.
      WithField("X-Road-Client", clientPathID).
      Error("Invalid X-Road-Client")

    respondWithError(w, http.StatusBadRequest, "Invalid X-Road-Client " + clientPathID)
    return
  }
  clientID := strings.ReplaceAll(clientPathID, "/", ":")

  providers, err := sph.sps.FindServiceProviders(clientID, serviceCode)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

 respondWithJSON(w, http.StatusOK, providers)
}

func (sh *serviceHandler) GetServiceServiceClients(w http.ResponseWriter, r *http.Request) {
  log.
    WithFields(log.Fields{
      "method": r.Method,
      "path":   r.URL,
    }).
    Info("Receive request")

  params := mux.Vars(r)
  serviceID := params["service-id"]

  clients, err := sh.scs.GetServiceClientsByService(serviceID)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

 respondWithJSON(w, http.StatusOK, clients)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
  respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)

  log.
    WithField("status_code", code).
    Info("Return response")

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}
