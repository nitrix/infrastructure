#!/bin/bash

# Deploy new stack.
docker stack deploy --prune -c docker-compose.yml prod

# Reload nginx configs.
docker ps -q -f name=prod_nginx | xargs -I % docker exec % nginx -s reload