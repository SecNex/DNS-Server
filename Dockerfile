FROM golang:latest AS builder

COPY ./src/go.mod ./src/go.sum ./

WORKDIR /go/src

RUN go mod download && go mod verify

# Copy the local package files to the container's workspace.
COPY ./src /go/src

# Build the outyet command inside the container.

ENV GOOS=linux
ENV GOARCH=amd64

# Path: Dockerfile
FROM debian:latest AS final

COPY --from=builder /go/bin/server /server
COPY ./config.json /config.json

EXPOSE 53/udp

RUN chmod +x /server

ENTRYPOINT ["/server"]