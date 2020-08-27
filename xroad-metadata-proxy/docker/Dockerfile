FROM golang:1.14-alpine AS develop
# clear GOPATH
ENV GOPATH= \
    GOCACHE=/tmp/.cache
RUN apk add --no-cache git make
COPY . /xroad-metadata-proxy
WORKDIR /xroad-metadata-proxy

FROM develop AS build
RUN make build

FROM alpine:3.9
COPY --from=build /xroad-metadata-proxy/bin/xroad-metadata-proxy_linux_amd64 ./xroad-metadata-proxy
EXPOSE 8083
CMD ["./xroad-metadata-proxy", "serve"]