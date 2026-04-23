FROM golang:1.24 AS builder

WORKDIR /src

COPY go.mod go.sum ./

ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY . .

# GOOS=linux GOARCH=amd64
RUN CGO_ENABLED=0 \
  go build -trimpath -ldflags="-s -w" \
  -o /out/transform-rpc ./services/rpc/transform






# FROM gcr.io/distroless/static-debian12:nonroot AS runtime
FROM alpine:latest AS runtime

WORKDIR /app
COPY --from=builder /out/transform-rpc ./transform-rpc
COPY services/rpc/transform/etc ./etc
EXPOSE 8081

ENTRYPOINT ["./transform-rpc", "-f", "etc/transform.docker.yaml"]








FROM golang:1.24 AS debug

WORKDIR /src

COPY go.mod go.sum ./

ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

# GOOS=linux GOARCH=amd64
RUN mkdir -p /app/etc && cp -r services/rpc/transform/etc/* /app/etc/ \
  && CGO_ENABLED=0 \
  go build -gcflags="all=-N -l" \
  -o /app/transform-rpc ./services/rpc/transform

WORKDIR /app
EXPOSE 8081 2345

ENTRYPOINT ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "--continue", "/app/transform-rpc", "--", "-f", "/app/etc/transform.docker.yaml"]
