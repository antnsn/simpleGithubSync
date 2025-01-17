# Use the official Golang image to build the Go application
FROM golang:1.18-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifest
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal image as the base image for the final stage
FROM alpine:latest

# Install git, openssh, and tzdata for timezone settings
RUN apk add --no-cache git openssh tzdata

# Set timezone using environment variable
ENV TZ=""
RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Create a directory for mounting volumes
RUN mkdir /repos

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /usr/local/bin/main

# Create the .ssh directory
RUN mkdir -p /root/.ssh

# Set permissions for the .ssh directory
RUN chmod 700 /root/.ssh

# Run the binary program produced by `go build`
ENTRYPOINT ["/usr/local/bin/main"]
