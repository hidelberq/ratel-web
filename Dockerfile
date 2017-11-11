FROM golang:1.9.1

MAINTAINER hidelberq <hide.seaweed@gmail.com>

RUN apt-get update

ADD . /go/src/github.com/hidelbreq/ratel-web
RUN go-wrapper download github.com/go-sql-driver/mysql
RUN go-wrapper install github.com/hidelbreq/ratel-web/server/...
WORKDIR /go/src/github.com/hidelbreq/ratel-web

CMD ["go", "run", "server/main.go"]
