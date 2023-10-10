#!/bin/bash

/chatait/backendServer/chatait-backend-server --gf.gcfg.path=/chatait/config &
/chatait/frontendServer/chatait-frontend-server --gf.gcfg.path=/chatait/config &
nginx -g "daemon off;"

while [[ true ]]; do
    sleep 1
done
