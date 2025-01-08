# # Base go image
# FROM golang:1.23-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o mailerApp ./cmd/api

# RUN chmod +x /app/mailerApp

# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY templates /templates

COPY mailerApp /app

CMD [ "/app/mailerApp" ]
