FROM golang:1.23.0-alpine as base

RUN apk add --no-cache gcc musl-dev git bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM base AS dev

RUN go install github.com/air-verse/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

FROM base AS build-production

EXPOSE 8080

CMD ["go", "run", "./cmd/app/main.go"]