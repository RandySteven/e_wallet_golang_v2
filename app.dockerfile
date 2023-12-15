FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go clean -mod=mod
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/rest/

EXPOSE 8080

CMD ./main