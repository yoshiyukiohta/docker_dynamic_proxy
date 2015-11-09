FROM ubuntu:14.04.3

MAINTAINER YoshiyukiOhta

RUN apt-get update
RUN apt-get install -y nginx-extras lua-nginx-redis
COPY nginx.conf /etc/nginx/nginx.conf
VOLUME ["/var/log/nginx"]
EXPOSE 80

CMD ["/usr/sbin/nginx"]
