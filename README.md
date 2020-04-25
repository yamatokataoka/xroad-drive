# File Service

File Service - a web service provides API interface and UI to share files over X-Road.

# Getting Started Development using Docker

On the project home directory, run

```
docker-compose --file docker-compose-dev.yml up --build
```

Access web UI via:

[http://localhost:8080](http://localhost:8080)

Remove docker containers and volumes:

```
docker-compose --file docker-compose-dev.yml down --volumes
```