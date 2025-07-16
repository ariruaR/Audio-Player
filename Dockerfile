# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev pkgconfig cairo-dev jpeg-dev giflib-dev

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o audio-player ./src

# Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache cairo jpeg giflib

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/audio-player .

# Create directory for images
RUN mkdir -p /app/src/image

# Copy image assets
COPY src/image/ /app/src/image/

# Expose port (if needed)
EXPOSE 8080

# Run the application
CMD ["./audio-player"]