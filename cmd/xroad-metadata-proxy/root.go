package main

import (
  "os"

  log "github.com/sirupsen/logrus"
  "github.com/spf13/cobra"
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
