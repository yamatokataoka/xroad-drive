# X-Road Drive

X-Road Drive - a web storage service to manage and share files over X-Road.

# Getting Started

On the project home directory, the docker-compose creates development environment.

```
./dockerized_dev.sh up -d --build
```

you can see logs

```
./dockerized_dev.sh logs -f
```

web UI http\://localhost:8080

API endpoints http\://localhost:8082



Remove docker containers and volumes

```
./dockerized_dev.sh down --volumes
```
