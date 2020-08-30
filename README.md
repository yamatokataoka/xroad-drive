# X-Road Drive

X-Road Drive - a web storage service to manage and share files over X-Road.

![screenshot](https://user-images.githubusercontent.com/34284486/91650666-44ab0a80-eabd-11ea-9c4d-3cf849f8e1e5.png)

## Getting Started

The docker compose creates the demo environment with demo data. On the project directory,

```
./dockerized_demo.sh up -d
```

web UI (client side) http\://localhost:8080

security server standalone UI http\://localhost:8084

Remove all docker containers and images.

```
./dockerized_demo.sh down --rmi all
```

## Docker Images
The Docker images are available on the Docker Hub.

* [yamatokataoka/xroad-drive-api](https://hub.docker.com/r/yamatokataoka/xroad-drive-api)
* [yamatokataoka/xroad-drive-ui](https://hub.docker.com/r/yamatokataoka/xroad-drive-ui)
* [yamatokataoka/xroad-metadata-proxy](https://hub.docker.com/r/yamatokataoka/xroad-metadata-proxy)

In the `docker-compose-demo.yml` and `docker-compose-dev.yml`, the pre-configured [niis/xroad-security-server-standalone](https://hub.docker.com/r/niis/xroad-security-server-standalone) is used.

* [yamatokataoka/xroad-drive-security-server-standalone-demo](https://hub.docker.com/r/yamatokataoka/xroad-drive-security-server-standalone-demo)
* [yamatokataoka/xroad-drive-security-server-standalone-dev](https://hub.docker.com/r/yamatokataoka/xroad-drive-security-server-standalone-dev)

## Development

The docker compose creates development environment.

On the project directory,

```
./dockerized_dev.sh up -d --build
```

web UI http\://localhost:8080

API http\://localhost:8082

metadata API http\://localhost:8083

security server standalone UI http\://localhost:8084

Remove docker containers and images.

```
./dockerized_dev.sh down --rmi all
```

## Contributing

Contributions are welcome!

1. Fork the repository.
1. Create a branch from develop with a descriptive name.
1. Implement your feature or bug fix.
1. Add, commit, and push your changes.
1. Submit a pull request.

## License
[MIT License](https://github.com/yamatokataoka/xroad-drive/blob/develop/LICENSE)
