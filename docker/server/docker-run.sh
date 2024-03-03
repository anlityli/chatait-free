#!/bin/sh

/chatait/backendServer/chatait-backend-server --gf.gcfg.path=/chatait/config &
/chatait/frontendServer/chatait-frontend-server --gf.gcfg.path=/chatait/config &

while [[ true ]]; do
    sleep 1
done
