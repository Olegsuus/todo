FROM golang:1.23.1-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .
COPY --from=builder /app/configs/ ./configs/
COPY --from=builder /app/migrations ./migrations

RUN apk add --no-cache curl bash postgresql-client

RUN set -ex; \
    ARCH=$(uname -m); \
    if [ "$ARCH" = "aarch64" ]; then \
      echo "Detected architecture: arm64"; \
      URL="https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_arm64"; \
    else \
      echo "Detected architecture: amd64"; \
      URL="https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_amd64"; \
    fi; \
    curl -L -o /usr/local/bin/goose "$URL"; \
    chmod +x /usr/local/bin/goose

EXPOSE 5555

CMD ["./app"]