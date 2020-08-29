OUTDIR := bin
VERSION := 0.1.0

serve:
	go run -ldflags "-X main.version=$(VERSION)" ./cmd/xroad-metadata-proxy serve

test:
	go test -race ./... -v

build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $(OUTDIR)/xroad-metadata-proxy_linux_amd64 cmd/xroad-metadata-proxy/*.go
