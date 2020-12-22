# Builder container
FROM golang:1.14.2 AS builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE on

WORKDIR /app

COPY . .
RUN go mod download && \
    go get -u github.com/cosmtrek/air

CMD ["air"]
