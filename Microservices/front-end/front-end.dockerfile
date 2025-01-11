# # Base go image
# FROM golang:1.23-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o frontApp ./cmd/api

# RUN chmod +x /app/frontApp

# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY frontApp /app

CMD [ "/app/frontApp" ]
