FROM golang:1.16-alpine

RUN mkdir /app

COPY brokerApp /app
CMD ["/app/brokerApp"]