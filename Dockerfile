FROM golang:1.22.2-alpine AS Dev
WORKDIR /root
COPY ./ ./
RUN go mod tidy
RUN go build -ldflags="-s -w" -o ./server ./cmd/main.go

FROM alpine:latest
WORKDIR /root
COPY --from=Dev /root/server /usr/local/bin/
EXPOSE 9092/tcp
ENTRYPOINT ["/usr/local/bin/server"]
CMD ["-l", "0.0.0.0:9092"]
