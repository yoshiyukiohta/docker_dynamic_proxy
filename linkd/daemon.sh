#!/bin/bash
sudo docker rm -f linkd

DOCKER_HOST=`ip addr show docker0 | grep -w inet | awk '{print $2;}' | cut -d"/" -f1`
UPSTREAMS_HOST=$(sudo docker inspect -f "{{ .NetworkSettings.IPAddress }}" upstreams)

sudo docker run -d \
    --name linkd \
    -v /var/run/docker.sock:/var/run/docker.sock \
    --link upstreams:redis \
    -e UPSTREAMS_HOST=tcp://$UPSTREAMS_HOST \
    sample/linkd
