package main

import (
  "net/http"

  log "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "github.com/gorilla/mux"

  proxy "github.com/yamatokataoka/xroad-metadata-proxy"
)

var serveCmd = &cobra.Command{
  Use:   "serve",
  Short: "",
  Long:  ``,
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 0 {
      cmd.Help()
    }
    serve()
  },
}

func serve() {
  var config proxy.Config

  err := viper.Unmarshal(&config)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to unmarshal config file")
  }

  router := mux.NewRouter().StrictSlash(true)
  client := proxy.NewClient(config.Database)
  defer client.Close()

  providerRepository := proxy.NewProviderRepository(client)
  clientRepository := proxy.NewClientRepository(client)
  providerService := proxy.NewProviderService(providerRepository)
  clientService := proxy.NewClientService(clientRepository)
  providerHandler := proxy.NewProviderHandler(providerService)
  clientHandler := proxy.NewClientHandler(clientService)

  router.HandleFunc("/service-providers", providerHandler.GetServiceProviders).Methods("GET")
  router.HandleFunc("/service-clients", clientHandler.GetServiceClients).Methods("GET")

  server := &http.Server{
    Handler: router,
    Addr:    config.Server.Addr,
  }

  log.WithFields(log.Fields{
    "address": server.Addr,
  }).Info("start listening")

  log.Fatal(server.ListenAndServe())
}

func init() {
  RootCmd.AddCommand(serveCmd)
}
