FROM golang:latest AS builder

WORKDIR /go/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./src/go.mod ./
COPY ./src/go.sum ./

RUN go mod download && go mod verify

COPY ./src .

RUN go build -v -o /tmp/app

# Path: Dockerfile
FROM ubuntu:latest

COPY --from=builder /tmp/app /usr/local/bin/app

COPY ./config.json /etc/config.json

ENTRYPOINT ["/usr/local/bin/app", "/etc/config.json"]