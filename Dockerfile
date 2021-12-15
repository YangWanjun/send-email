FROM golang:1.17.5-alpine3.15

WORKDIR /opt/project
COPY go.mod /opt/project/go.mod
COPY go.sum /opt/project/go.sum

RUN go mod download

EXPOSE 80