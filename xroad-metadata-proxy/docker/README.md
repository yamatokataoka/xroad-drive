# Quick reference
* **Github Repository**: [the X-Road Drive](https://github.com/yamatokataoka/xroad-drive)

## Tags
* develop
* latest
* 0.0.1
* 0.1.0

## Environment Variables
### `XROAD_METADATA_PROXY_DATABASE_ADDR`

This variable specifies the TCP address for the database to listen on, in the form "host:port". The default is `localhost:6379`.

# What is xroad-metadata-proxy?
The xroad-metadata-proxy is a proxy for metadata service and admin api on X-Road.

It is developed as a part of the [xroad-drive](https://github.com/yamatokataoka/xroad-drive).

See [X-Road Data Exchange Layer](https://github.com/nordic-institute/X-Road)

# Usage
The xroad-metadata-proxy docker container requires Redis.

Use the following docker-compose file as an example
```
services:
  xroad-metadata-proxy:
    image: yamatokataoka/xroad-metadata-proxy:latest
    ports:
      - 8083:8083
    environment:
      - XROAD_METADATA_PROXY_DATABASE_ADDR=redis:6379
    depends_on:
      - redis
  redis:
    image: redis:latest
    ports:
      - 6379:6379
```

You can find the full demo docker compose file, [docker-compose-demo.yml](https://github.com/yamatokataoka/xroad-drive/blob/master/docker-compose-demo.yml).
