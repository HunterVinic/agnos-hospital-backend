# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed for private or indirect dependencies)
RUN apk add --no-cache git

# Copy go mod files first (better Docker caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd

# ---------- Runtime Stage ----------
FROM alpine:latest

WORKDIR /app

# Install CA certificates (important for HTTPS calls)
RUN apk add --no-cache ca-certificates

# Copy compiled binary from builder
COPY --from=builder /app/main .

# Expose application port
EXPOSE 8080

# Run the app
CMD ["./main"]