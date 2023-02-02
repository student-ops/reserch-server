FROM golang:1.18-alpine3.16 AS builder

RUN mkdir -p /app/api
WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY ./cmd/api/*.go ./cmd/api/
RUN mkdir  ./db-manage
COPY ./db-manage ./db-manage

RUN go build -o surroundings-apiApp /app/cmd/api/*.go

FROM alpine:3.16

COPY --from=builder /app/surroundings-apiApp /
CMD ["/surroundings-apiApp"]