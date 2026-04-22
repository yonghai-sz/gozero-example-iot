FROM golang:1.26 AS builder

WORKDIR /src

COPY go.mod go.sum ./

ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -trimpath -ldflags="-s -w" \
  -o /out/shorturl-api ./api









# FROM gcr.io/distroless/static-debian12:nonroot AS runtime
FROM alpine:latest AS runtime

# USER nobody

WORKDIR /app
COPY --from=builder /out/shorturl-api ./shorturl-api
COPY api/etc ./etc
EXPOSE 8080

ENTRYPOINT ["./shorturl-api", "-f", "etc/shorturl-api.docker.yaml"]









FROM golang:1.26 AS debug

WORKDIR /src

COPY go.mod go.sum ./

ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .

RUN mkdir -p /app/etc \
  && cp -r api/etc/* /app/etc/ \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -gcflags="all=-N -l" \
  -o /app/shorturl-api ./api

WORKDIR /app
EXPOSE 8080 2345

ENTRYPOINT ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/shorturl-api", "--", "-f", "/app/etc/shorturl-api.docker.yaml"]
