OUTDIR := bin

serve:
	go run ./cmd/xroad-metadata-proxy serve

test:
	go test -race ./... -v

build:
	GOOS=linux GOARCH=amd64 go build -o $(OUTDIR)/xroad-metadata-proxy_linux_amd64 cmd/xroad-metadata-proxy/*.go