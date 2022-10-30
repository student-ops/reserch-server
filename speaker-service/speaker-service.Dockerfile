FROM golang:1.18-alpine3.16 AS builder

# definatery need to be separated
RUN mkdir -p  /app/cmd/api/data
WORKDIR /app

ENV GO111MODULE=on
COPY go.mod go.sum /app/
COPY cmd/api/*.go /app/cmd/api/
COPY cmd/api/data/*.go /app/cmd/api/data/
RUN go mod download
# RUN go get main/data
# RUN go install
RUN go build -o speakerApp /app/cmd/api/*.go

FROM alpine:3.16

COPY --from=builder /app/speakerApp /
CMD ["/speakerApp"]