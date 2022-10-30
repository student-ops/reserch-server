FROM golang:1.18-alpine3.16 AS builder

RUN mkdir /tool

WORKDIR /tool

COPY ./go.mod .
COPY ./go.sum .
COPY ./cmd/api/*.go .
RUN go build -o brokertool ./*.go

FROM alpine:3.16

COPY --from=builder /tool/brokertool .
RUN mkdir /sample
COPY sample /sample
CMD ["/brokertool"]