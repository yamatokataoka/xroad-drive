package proxy

import (
  "net/http"
  "encoding/json"

  log "github.com/sirupsen/logrus"
)

type ProviderHandler interface {
  GetServiceProviders(w http.ResponseWriter, r *http.Request)
}

type ClientHandler interface {
  GetServiceClients(w http.ResponseWriter, r *http.Request)
}

type providerHandler struct {
  ps ProviderService
}

type clientHandler struct {
  cs ClientService
}

func NewProviderHandler(ps ProviderService) ProviderHandler {
  return &providerHandler{ps}
}

func NewClientHandler(cs ClientService) ClientHandler {
  return &clientHandler{cs}
}

func (ph *providerHandler) GetServiceProviders(w http.ResponseWriter, r *http.Request) {
  log.
    WithFields(log.Fields{
      "method":     r.Method,
      "path":       r.URL,
    }).
    Info("Receive request")

  providers, err := ph.ps.GetAll()
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

 respondWithJSON(w, http.StatusOK, providers)
}

func (ch *clientHandler) GetServiceClients(w http.ResponseWriter, r *http.Request) {
  log.
    WithFields(log.Fields{
      "method":     r.Method,
      "path":       r.URL,
    }).
    Info("Receive request")

  clients, err := ch.cs.GetAll()
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
