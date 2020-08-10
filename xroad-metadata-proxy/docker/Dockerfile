FROM golang:1.14-alpine AS develop
# clear GOPATH
ENV GOPATH= \
    GOCACHE=/tmp/.cache
RUN apk add --no-cache git make
COPY . /xroad-metadata-proxy
WORKDIR /xroad-metadata-proxy

FROM develop AS build
RUN go build -o ./out/xroad-metadata-proxy .

FROM alpine:3.9
COPY --from=build ./out/xroad-metadata-proxy /xroad-metadata-proxy
EXPOSE 8083
CMD ["/xroad-metadata-proxy"]