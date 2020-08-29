#!/bin/bash
sleep 2
echo 'Starting seeding ...'
redis-cli < /docker-entrypoint-initdb.d/init.redis
echo 'Finished seeding ...'
