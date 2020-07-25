package main

import (
  "net/http"

  log "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
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
  router := mux.NewRouter().StrictSlash(true)
  client := proxy.NewClient()

  providerRepository := proxy.NewProviderRepository(client)
  providerService := proxy.NewProviderService(providerRepository)
  providerHandler := proxy.NewProviderHandler(providerService)

  router.HandleFunc("/service-providers", providerHandler.GetServiceProviders).Methods("GET")

  server := &http.Server{
    Handler: router,
    Addr:    "127.0.0.1:8083",
  }

  log.WithFields(log.Fields{
    "address": server.Addr,
  }).Info("start listening")

  log.Fatal(server.ListenAndServe())
}

func init() {
  RootCmd.AddCommand(serveCmd)
}
