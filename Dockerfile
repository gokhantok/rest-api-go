# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.11 base image
FROM golang:alpine as build-env

# Add Maintainer Info
LABEL maintainer="Gökhan Tok <gokhan.tok@outlook.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/go-docker

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container

COPY . $GOPATH/src/go-docker

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the app

RUN go build -o myapp ./cmd/myapp/*.go
EXPOSE 8000 8000


# final stage
FROM alpine
# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/go-docker
COPY --from=build-env  go/src/go-docker/myapp .

EXPOSE 8000 8000
# Run the executable 
CMD ["./myapp"]