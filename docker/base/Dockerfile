FROM golang:1.21-alpine AS builder

WORKDIR /source

ENV CGO_ENABLED 0
ENV GOOS linux

COPY . .

RUN go mod download \
    && go mod tidy \
    && go build -o bin/api cmd/api/main.go \
    && go build -o bin/url cmd/url/main.go \
