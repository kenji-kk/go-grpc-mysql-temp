FROM golang:latest

WORKDIR /go/api

ADD ./api /go/api

RUN go mod tidy
