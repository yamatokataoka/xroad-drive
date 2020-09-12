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
  log.
    WithField("version", version).
    Info("xroad-metadata-proxy")

  var config proxy.Config

  err := viper.Unmarshal(&config)
  if err != nil {
    log.
      WithError(err).
      Error("Failed to unmarshal config file")
  }

  router := mux.NewRouter().StrictSlash(true)

  metaClient := proxy.NewMetadataServiceClient(config.SecurityServer)
  adminClient := proxy.NewAdminAPIClient(config.SecurityServer)

  srp := proxy.NewServiceProviderRepository(metaClient)
  sr := proxy.NewServiceRepository(metaClient)
  scr := proxy.NewServiceClientRepository(adminClient)

  sps := proxy.NewServiceProviderService(srp, sr)
  scs := proxy.NewServiceClientService(scr)

  sph := proxy.NewServiceProviderHandler(sps)
  sh := proxy.NewServiceHandler(scs)

  router.
    HandleFunc("/service-providers", sph.GetServiceProviders).
    Methods("GET")
  router.
    HandleFunc("/services/{service-id}/service-clients", sh.GetServiceServiceClients).
    Methods("GET")

  server := &http.Server{
    Handler: router,
    Addr:    config.Server.Addr,
  }

  log.
    WithField("address", server.Addr).
    Info("Start listening")

  log.Fatal(server.ListenAndServe())
}

func init() {
  RootCmd.AddCommand(serveCmd)
}
