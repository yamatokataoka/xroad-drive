# Quick reference
* **the X-Road Drive Github Repository**: [xroad-drive](https://github.com/yamatokataoka/file-service/tree/develop)

## Tags
* develop
* latest
* 1.0.0

## Environment Variables
You can use Spring boot common application properties as environment veriables to configure the mongodb connection setting.

For example,
```
services:
  xroad-drive-api:
    image: yamatokataoka/xroad-drive-api:latest
    environment:
      - SPRING_DATA_MONGODB_HOST=mongodb
```

You can find more properties, [here](https://docs.spring.io/spring-boot/docs/current/reference/html/appendix-application-properties.html#data-properties).

# What is xroad-drive-api?
The xroad-drive-api is a REST api to manage and share files over X-Road.

about [X-Road Data Exchange Layer](https://github.com/nordic-institute/X-Road)

# Usage
The xroad-drive-api docker container requires a mongodb.

Use following docker-compose file as a example
```
services:
  xroad-drive-api:
    image: yamatokataoka/xroad-drive-api:latest
    ports:
      - 8082:8082
    environment:
      - SPRING_DATA_MONGODB_HOST=mongodb
    depends_on:
      - mongodb
  mongodb:
    image: mongo:latest
    ports:
      - 27017:27017
```

You can find the full demo docker compose file, [docker-compose-demo.yml](https://github.com/yamatokataoka/xroad-drive/blob/master/docker-compose-demo.yml).

# Where to Store Files

**Important note**: There are several ways to store files uploaded by users.

Write docker-compose file like
```
services:
  xroad-drive-api:
    image: yamatokataoka/xroad-drive-api:latest
    volumes:
      - /my/own/datadir:/etc/xroad-drive/xroad-drive-api/files
```

The `- /my/own/datadir:/etc/xroad-drive/xroad-drive-api/files` part of the file mounts the `/my/own/datadir` directory from the underlying host system as `/etc/xroad-drive/xroad-drive-api/files` inside the container, where xroad-drive-api by default will write its file data.
