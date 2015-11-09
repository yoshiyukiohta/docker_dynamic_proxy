#!/bin/bash

sudo docker rm -f upstreams

sudo docker run -d -P --name upstreams redis
