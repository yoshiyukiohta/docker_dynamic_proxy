#!/bin/bash
sudo docker rm -f dproxy
sudo docker run -d -p 80:80 -v /var/log/nginx:/var/log/nginx --name dproxy --link upstreams:redis sample/dproxy
