FROM golang:1.17-buster

RUN apt-get update
RUN apt-get install -y protobuf-compiler

ENV PATH="$PATH:$(go env GOPATH)/bin"

EXPOSE 80
