FROM golang:1.9.1

MAINTAINER hidelberq <hide.seaweed@gmail.com>

ADD . /go/src/github.com/hidelbreq/ratel-web
RUN go get -u github.com/go-sql-driver/mysql
RUN go install github.com/hidelbreq/ratel-web/server/...
WORKDIR /go/bin

EXPOSE 8080

