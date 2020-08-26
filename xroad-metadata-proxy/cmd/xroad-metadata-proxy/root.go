package main

import (
  "os"

  log "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"

  proxy "github.com/yamatokataoka/xroad-metadata-proxy"
)

var RootCmd = &cobra.Command{
  Use:   "xroad-metadata-proxy",
  Short: "X-Road metadata proxy service.",
  Long: ``,
}

func Execute() {
  if err := RootCmd.Execute(); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(proxy.InitConfig)
}
