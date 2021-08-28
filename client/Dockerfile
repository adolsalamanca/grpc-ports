FROM golang:alpine as builder

COPY . /src
WORKDIR /src

RUN go build -o client cmd/main.go

FROM alpine:latest as runner

COPY --from=builder /src/client /opt/client/
WORKDIR /opt/client

ENTRYPOINT ["./client", "-grpc_server_host=grpc-ports_server_1"]

