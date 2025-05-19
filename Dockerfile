FROM golang:1.23

WORKDIR /app

COPY ./api-gateway/go.mod ./api-gateway/go.sum ./

COPY ./user-service ./user-service
COPY ./car-service ./car-service
COPY ./rental-service ./rental-service
COPY ./api-gateway ./api-gateway

WORKDIR /app/api-gateway

RUN go mod tidy
RUN go build -o main ./cmd/main.go

CMD ["./main"]
