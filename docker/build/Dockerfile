FROM golang:1.22-alpine

ENV USER=dev

WORKDIR /app

## Dependências
# RUN apk add

## Create a group and user
RUN addgroup -g 1000 dev_group
RUN adduser -u 1000 -G dev_group -D dev
