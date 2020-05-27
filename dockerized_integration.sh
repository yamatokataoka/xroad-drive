#!/usr/bin/env bash
set -ex

function docker-compose-base () {
  docker-compose --file docker-compose-integration-with-xroad.yml "$@"
}

docker-compose-base $*

if [[ "$*" == *up* ]]; then
  docker-compose-base
    exec --user root xroad-drive-api \
    chown -R xroad-drive-api-user:xroad-drive-api-user /etc/xroad-drive/xroad-drive-api/files
  
  docker-compose-base exec xroad-security-server-standalone /bin/bash -c "
      cp -pn /etc/xroad/nginx/default-xroad.conf /etc/xroad/nginx/default-xroad.conf.backup &&
      sed "3,17s/^/#/" /etc/xroad/nginx/default-xroad.conf.backup > /etc/xroad/nginx/default-xroad.conf &&
      nginx -s reload
    "
fi