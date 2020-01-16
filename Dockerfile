FROM golang:1.13.6-alpine3.11

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work
