FROM --platform=arm64 golang:latest AS builder

# Copy the local package files to the container's workspace.
COPY ./src /go/src

WORKDIR /go/src

# Build the outyet command inside the container.
RUN go install

ENV GOOS=linux
ENV GOARCH=arm64

RUN GOOS=linux GOARCH=arm64 go build -o /go/bin/server

# Path: Dockerfile
FROM --platform=arm64 debian:latest AS final

COPY --from=builder /go/bin/server /server
COPY ./config.json /config.json

EXPOSE 53/udp

RUN chmod +x /server

ENTRYPOINT ["/server"]