FROM golang:1.18-alpine3.16 AS builder

RUN mkdir /app

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY ./cmd/api/*.go .
RUN go build -o brokerApp ./*.go

FROM alpine:3.16

COPY --from=builder /app/brokerApp .
RUN mkdir /sample
COPY sample /sample
CMD ["/brokerApp"]