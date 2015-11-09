#!/bin/bash

UPSTREAMS_HOST=$(sudo docker inspect -f "{{ .NetworkSettings.IPAddress }}" dproxy_upstreams_1)

sudo docker run -it --rm \
    redis \
    redis-cli \
    -h ${UPSTREAMS_HOST:?}
