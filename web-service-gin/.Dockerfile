# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build

EXPOSE 4001

CMD ["./produce-api"]