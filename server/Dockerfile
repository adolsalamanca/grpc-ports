FROM golang:alpine as builder

COPY . /src
WORKDIR /src

RUN go build -o server app/main.go

FROM alpine:latest as runner

COPY --from=builder /src/server /opt/server/
WORKDIR /opt/server

ENTRYPOINT ["./server"]

