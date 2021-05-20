# 開発用
FROM golang:1.15.7-alpine as dev

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download
EXPOSE 8080