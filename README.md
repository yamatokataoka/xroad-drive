# File Service

File Service - a web strage service provides API interface and UI to share files over X-Road.

# Getting Started Development using Docker

On the project home directory, run

```
./dockerized_dev.sh up --build
```

Access web UI via:

[http://localhost:8080](http://localhost:8080)

Remove docker containers and volumes:

```
./dockerized_dev.sh down --volumes
```