FROM golang:1.18-alpine3.16 AS builder

RUN mkdir -p /app/cmd
WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY ./cmd/*.go ./cmd

RUN go build -o reminder-serviceApp /app/cmd/*.go

FROM alpine:3.16

COPY --from=builder /app/reminder-serviceApp /
CMD ["/reminder-serviceApp"]