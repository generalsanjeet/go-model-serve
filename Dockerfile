# Stage 1: Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first (to leverage Docker layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Now copy the rest
COPY . .

# Build the binary
RUN go build -o server ./cmd/server

# Stage 2: Runtime image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Set environment variable to ensure logs are shown
ENV GIN_MODE=release

CMD ["./server"]

