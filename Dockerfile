# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary
COPY --from=builder /app/main .

# Copy frontend files
COPY --from=builder /app/frontend/dist ./frontend/dist

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./main"]
