# 開発用
FROM golang:1.15.7-alpine as dev

WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 9000