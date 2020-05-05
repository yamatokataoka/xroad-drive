#!/usr/bin/env bash
set -ex

export uid=$(id -u)
export gid=$(id -g)

docker-compose --file docker-compose-dev.yml $*
