# Stage 1: The Build Stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# Build the Go application binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-s -w' -o kamal .

# Stage 2: The Final Image
FROM alpine:latest

WORKDIR /

# Copy the compiled binary from the builder stage
COPY --from=builder /app/kamal .

# Set executable permissions for the binary
RUN chmod +x kamal

# Expose the port your application listens on
EXPOSE 3000

# The command to run when the container starts
CMD ["./kamal"]
