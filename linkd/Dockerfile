FROM golang:1.5.1

MAINTAINER YoshiyukiOhta

RUN go get github.com/fsouza/go-dockerclient
RUN go get github.com/garyburd/redigo/redis

COPY linkd.go ./linkd.go

CMD ["go", "run", "linkd.go"]
