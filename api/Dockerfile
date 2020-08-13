# Start from golang base image
#FROM golang:alpine as builder
FROM golang:1.14.2-alpine as builder

# ENV GO111MODULE=on

# Add Image info
LABEL name='Golang 1.14 (alpine) Image'
LABEL maintainer="Damien Bernard <damien.bernard@epitech.eu>"
LABEL version="1.0"
LABEL description="Temporary Image for Multi stage usage in Cbm"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update                                          \
    && apk add --no-cache git                           \
    && apk add gcc                                      \
    && apk add libc-dev

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
LABEL name='Cbm - go Rest API'
LABEL maintainer="Damien Bernard <damien.bernard@epitech.eu>"
LABEL version="1.0"
LABEL description="Go Api for cbm"

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080
#EXPOSE $API_PORT

#Command to run the executable
#CMD ["./main"]
ENTRYPOINT ./main
