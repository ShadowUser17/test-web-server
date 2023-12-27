#### Dependencies:
- [gin](https://github.com/gin-gonic/gin/tree/v1.9.1)
- [prometheus](https://github.com/prometheus/client_golang/tree/v1.17.0)

#### Build binary file:
```bash
go mod tidy
```
```bash
go build -ldflags="-s -w" -o ./server ./cmd/main.go
```

#### Build docker image:
```bash
docker build -t "shadowuser17/test-web-server:latest" .
```

#### Publish docker image:
```bash
docker login -u "${DOCKERHUB_LOGIN}" -p "${DOCKERHUB_TOKEN}"
```
```bash
docker push --all-tags "shadowuser17/test-web-server"
```
