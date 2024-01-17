FROM golang:1.21-alpine

WORKDIR /app
COPY ./lib ./lib
COPY ./src/load-balancer ./src/load-balancer
COPY ./config/load-balancer.json ./config/load-balancer.json

WORKDIR /app/src/load-balancer
RUN go build -o ./main

CMD ["/app/src/load-balancer/main"]