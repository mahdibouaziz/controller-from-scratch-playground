# First stage: Build the Go binary
FROM golang:1.23 AS builder

# Set working directory inside the container
WORKDIR /workspace

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the manager binary (static binary for minimal runtime compatibility)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /manager ./cmd/main.go

# Second stage: Use Alpine as the minimal runtime environment
FROM alpine:latest

WORKDIR /

# Install ca-certificates (needed for HTTPS calls)
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /manager /manager

# Run the manager binary
ENTRYPOINT ["/manager"]
