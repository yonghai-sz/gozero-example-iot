FROM golang:1.24 AS builder
WORKDIR /src
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# GOOS=linux GOARCH=amd64 
RUN CGO_ENABLED=0 \
  go build -trimpath -ldflags="-s -w" -o /out/shorturl-api ./services/api




# FROM gcr.io/distroless/static-debian12:nonroot AS runtime
FROM alpine:latest AS runtime
# USER nobody
WORKDIR /app
COPY --from=builder /out/shorturl-api ./shorturl-api
COPY services/api/etc ./etc
EXPOSE 8080
ENTRYPOINT ["./shorturl-api", "-f", "etc/shorturl-api.docker.yaml"]









   











FROM golang:1.24 AS dlv
WORKDIR /src
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/go-delve/delve/cmd/dlv@latest

FROM builder AS debug-builder
RUN CGO_ENABLED=0 \
  go build -gcflags="all=-N -l" -o /out/shorturl-api ./services/api
  
FROM golang:1.24 AS debug
WORKDIR /app
COPY --from=debug-builder /out/shorturl-api ./shorturl-api
COPY --from=dlv /go/bin/dlv /usr/local/bin/dlv
COPY services/api/etc ./etc
EXPOSE 8080 2345
ENTRYPOINT ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "--continue", "/app/shorturl-api", "--", "-f", "/app/etc/shorturl-api.docker.yaml"]
