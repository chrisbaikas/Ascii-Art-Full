# Use the official minimal Go image based on Alpine Linux
FROM golang:1.22-alpine

# Install bash to allow usage of /bin/bash inside the container
RUN apk add --no-cache bash

# Add metadata information to the image
LABEL maintainer="chbaikas, askordal, npanouts"
LABEL version="1.0"
LABEL description="ASCII Art Web Server in Go"

# Set the working directory inside the container
WORKDIR /app

# Copy all files from the current directory to the container
COPY . .

# Build the Go application and name the binary ascii-art-web
RUN go build -o ascii-art-web main.go

# Expose port 8080 for incoming HTTP traffic
EXPOSE 8080

# Define the command to run when the container starts
CMD ["./ascii-art-web"]
