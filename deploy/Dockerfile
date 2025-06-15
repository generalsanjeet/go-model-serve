# ---------- Stage 1: Build ----------
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go server binary
RUN go build -o server ./cmd/server


# ---------- Stage 2: Runtime ----------
FROM gcr.io/distroless/static:nonroot

WORKDIR /app

# Copy the compiled binary from builder stage
COPY --from=builder /app/server .

# Run as non-root for security
USER nonroot:nonroot

# Start the server
ENTRYPOINT ["./server"]

