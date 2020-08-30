#!/bin/bash
sleep 2
echo 'Start seeding ...'
mongorestore --drop \
  --db test \
  --collection metadata \
  /docker-entrypoint-initdb.d/init.bson
echo 'Finish seeding ...'
