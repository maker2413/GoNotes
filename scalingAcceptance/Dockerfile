# Make sure to specify the same Go version as the one in the go.mod file.
# For example, golang:1.22.1-alpine.
FROM golang:1.23.2-alpine

WORKDIR /app

ARG bin_to_build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/${bin_to_build}/main.go

CMD [ "./svr" ]
