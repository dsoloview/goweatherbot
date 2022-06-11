FROM golang:latest

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY ./ /app

RUN apt-get install git
RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon --build="go build cmd/bot/main.go" --command=./main

