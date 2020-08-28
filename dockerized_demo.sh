#!/usr/bin/env bash
set -ex

docker-compose \
  --file docker-compose-demo.yml \
  --project-name xroad-drive-demo $*
