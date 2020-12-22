# Builder container
FROM golang:1.14.2 AS builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE on

WORKDIR /go/src/github.com/yamalabo

COPY . .
RUN go mod download && \
    go get github.com/markbates/refresh


RUN go build -o main

EXPOSE 8084

CMD ["./main"]

