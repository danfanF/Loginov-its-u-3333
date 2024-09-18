FROM golang:1.21-alpine as builder

WORKDIR /service

RUN apk --no-cache add bash git make gcc gettext musl-dev openssh-client

COPY . .

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o /app/main ./cmd

FROM alpine:latest

COPY --from=builder /app/main /app/main

#ENTRYPOINT ["/app/main"]