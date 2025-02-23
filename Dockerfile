ARG GO_VERSION=1.22.2

# Builder stage
FROM golang:${GO_VERSION} AS builder
WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o dist/api cmd/api/main.go
RUN go build -o dist/worker cmd/worker/main.go

FROM debian:12.5-slim
WORKDIR /app

RUN apt-get update && \
    apt-get install -y ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/dist/ .
COPY --from=builder /app/pkg/db/migration ./pkg/db/migration

RUN chmod +x /app/api /app/worker

ENV PORT=12345
EXPOSE $PORT
ENV GIN_MODE=release

CMD ["sh", "-c", "trap 'kill 0' SIGINT SIGTERM; /app/api & /app/worker & wait"]
