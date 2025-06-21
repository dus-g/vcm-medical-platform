FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy the binary
COPY --from=builder /app/main .

# Copy frontend build files if they exist
COPY --from=builder /app/frontend/dist ./frontend/dist

EXPOSE 8080
CMD ["./main"]
