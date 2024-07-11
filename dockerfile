# Use the official Golang image to build the Go application
FROM golang:1.18-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal image as the base image for the final stage
FROM alpine:latest

# Install git and openssh
RUN apk add --no-cache git openssh

# Set environment variables for directories and SSH key
ENV FOLDER_PATHS="" \
    SSH_KEY=""

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /usr/local/bin/main

# Run the binary program produced by `go build`
ENTRYPOINT ["/usr/local/bin/main"]
