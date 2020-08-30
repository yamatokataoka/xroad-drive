package proxy

import (
  "strings"

  log "github.com/sirupsen/logrus"
  "github.com/spf13/viper"
)

type Config struct {
  Server       ServerConfig
  Database     DatabaseConfig
}

type ServerConfig struct {
	Addr string
}

type DatabaseConfig struct {
	Addr string
}

func InitConfig() {
  viper.SetDefault("server.addr", "0.0.0.0:8083")
  viper.SetDefault("database.addr", "localhost:6379")

  viper.SetConfigName("xroad-metadata-proxy")
  viper.AddConfigPath(".")

  viper.AutomaticEnv()
  viper.SetEnvPrefix("XROAD_METADATA_PROXY")
  viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

  err := viper.ReadInConfig()
  if err != nil {
    log.
      WithError(err).
      Error("Failed to read config file")
  }

  // Workaround because viper does not treat env vars the same as other config
  viper.BindEnv("database.addr")
  viper.BindEnv("server.addr")
}
