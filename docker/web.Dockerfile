# Stage 1: Build Go Application
FROM golang:1.23-alpine as base

# Install dependencies for building Go apps and FFmpeg
RUN apk add --no-cache build-base git inotify-tools bash

# Install sql-migrate
RUN go install github.com/rubenv/sql-migrate/...@latest

# Set working directory
WORKDIR /clean_web

# Copy go.mod and go.sum separately to leverage Docker cache
COPY go.mod go.sum ./

# Initialize Go modules
RUN go mod download

# Copy all files from your repo to the container
COPY . .

# Build Go application
RUN go build -o main .

# Clean up build dependencies to reduce image size
RUN apk del build-base git && rm -rf /var/cache/apk/*


# Stage 2: Final Image with FFmpeg
FROM jrottenberg/ffmpeg:4.4-alpine

# Install necessary runtime dependencies
RUN apk add --no-cache bash inotify-tools git

# Download and install Go version 1.23 manually
RUN apk add --no-cache --virtual .build-deps curl && \
    curl -fsSL https://golang.org/dl/go1.23.0.linux-amd64.tar.gz | tar -C /usr/local -xz && \
    apk del .build-deps

# Set Go environment variables
ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$PATH

# Install sql-migrate
RUN go install github.com/rubenv/sql-migrate/...@latest

# Copy the built Go binary from the first stage
COPY --from=base /go/bin/sql-migrate /usr/local/bin/
COPY --from=base /clean_web/main /usr/local/bin/

# Set working directory for your final app
WORKDIR /clean_web

# Copy all files from your repo to the final image (including run.sh and other necessary files)
COPY . .

# Copy the script to run the application
COPY docker/run.sh /clean_web/docker/

# Default entrypoint
ENTRYPOINT ["/bin/sh", "/clean_web/docker/run.sh"]
