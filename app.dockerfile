FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go clean -mod=mod
RUN go mod download && go mod verify

