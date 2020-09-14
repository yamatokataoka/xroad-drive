package proxy

import (
  "time"
  "net/http"
  "crypto/tls"
)

type HTTPClient interface {
  Do(req *http.Request) (*http.Response, error)
}

type MetadataServiceClient struct {
  BaseURL    string
  HTTPClient HTTPClient
}

type AdminAPIClient struct {
  BaseURL    string
  apiKey     string
  HTTPClient HTTPClient
}

func NewMetadataServiceClient(config SSConfig) *MetadataServiceClient {
  return &MetadataServiceClient{
    BaseURL: config.URL,
    HTTPClient: &http.Client{
      Timeout: time.Minute,
    },
  }
}

func NewAdminAPIClient(config SSConfig) *AdminAPIClient {
  return &AdminAPIClient{
    BaseURL: config.AdminURL,
    apiKey:  config.APIKey,
    HTTPClient: &http.Client{
      Timeout: time.Minute,
      Transport: &http.Transport{
        TLSClientConfig: &tls.Config{
          InsecureSkipVerify: true,
        },
      },
    },
  }
}
