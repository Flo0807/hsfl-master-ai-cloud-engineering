FROM golang:1.21-alpine

WORKDIR /app
COPY ./lib ./lib
COPY ./src/api-gateway ./src/api-gateway
COPY ./config/api-gateway.json ./config/api-gateway.json

WORKDIR /app/src/api-gateway
RUN go build -o ./main

CMD ["/app/src/api-gateway/main"]