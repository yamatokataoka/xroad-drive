#!/usr/bin/env bash
set -ex

export LOCAL_UID=$(id -u)
export LOCAL_GID=$(id -g)

docker-compose --file docker-compose-dev.yml $*
