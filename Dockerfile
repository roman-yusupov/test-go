# syntax=docker/dockerfile:1

FROM golang:1.17.1
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go   ./
COPY ./proto/ ./proto/

RUN go mod download

RUN go build -o /server
EXPOSE 5100
CMD [ "/server" ]

