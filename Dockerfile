# syntax=docker/dockerfile:1
# Start from golang base image
FROM golang:alpine as builder

ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Igor Ariza <igorariza@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:3.4
RUN apk --no-cache add ca-certificates
RUN apk update && apk add bash
RUN /bin/sh -c "apk add --no-cache bash"
RUN apk --no-cache add curl jq
RUN apk --no-cache add mysql mysql-client

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 8000

COPY entrypoint.sh ./entrypoint.sh

#Command to run the executable
#ENTRYPOINT [ "/bin/sh", "./entrypoint.sh" ]
CMD ["./main"]