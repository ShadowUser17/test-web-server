VERSION 0.8
FROM golang:1.23.9-alpine
WORKDIR /root

build:
    COPY cmd cmd
    COPY go.* .
    RUN go mod tidy
    RUN go build -ldflags="-s -w" -o server cmd/main.go
    SAVE ARTIFACT server

docker:
    ARG tag="latest"
    FROM alpine:latest
    COPY +build/server /usr/local/bin/
    EXPOSE 9092/tcp
    ENTRYPOINT ["/usr/local/bin/server"]
    CMD ["-l", "0.0.0.0:9092"]
    SAVE IMAGE --push "shadowuser17/test-web-server:$tag"

all:
    BUILD +build
    BUILD +docker
