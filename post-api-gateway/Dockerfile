FROM golang:1.17-buster

RUN apt-get update
RUN apt-get install -y protobuf-compiler

# RUN go get -u github.com/gin-gonic/gin
# RUN go get -u golang.org/x/crypto
# RUN go get github.com/google/uuid
# RUN go get github.com/go-redis/redis/v8
# RUN go get -u google.golang.org/protobuf
# RUN go get -u google.golang.org/protobuf/proto
# # RUN go get -u google.golang.org/protobuf/protoc-gen-go
# RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# RUN go get google.golang.org/grpc

ENV PATH="$PATH:$(go env GOPATH)/bin"

EXPOSE 80

# FROM        golang
# RUN         mkdir -p /app
# WORKDIR     /app
# COPY        . .
# RUN         go mod download
# RUN         go build -o app
# ENTRYPOINT  ["./app"]
