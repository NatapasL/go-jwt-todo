ARG go_version=1.15

FROM golang:${go_version}-alpine AS builder

ENV GIN_MODE=release
ENV PORT=3000
ENV APP_PATH=/go/src/github.com/NatapasL/go-jwt-todo

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./app ./main.go

EXPOSE $PORT

ENTRYPOINT ["./app"]
