# First stage: Build the Go binary
FROM golang:1.23 AS builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build a fully static binary
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /manager ./cmd/main.go
RUN go build -ldflags="-s -w" -o /manager ./cmd/main.go

# # Second stage: Use Debian Slim as the runtime
FROM debian:bookworm-slim

WORKDIR /

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /manager /manager

ENTRYPOINT ["/manager"]
