FROM golang:1.20.8-alpine3.18
RUN apk update && apk add git
RUN mkdir /go/app
WORKDIR /go/app
ADD ./ /go/app

ENV GO111MODULE=on
# RUN go install github.com/cosmtrek/air@latest

RUN go mod download