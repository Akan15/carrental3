FROM golang:1.24.1


# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum проекта
COPY ./api-gateway/go.mod ./api-gateway/go.sum ./

# Копируем все нужные папки микросервисов и api-gateway
COPY ./user-service ./user-service
COPY ./car-service ./car-service
COPY ./rental-service ./rental-service
COPY ./api-gateway ./api-gateway

# Переходим в директорию api-gateway
WORKDIR /app/api-gateway

# Загружаем зависимости и билдим
RUN go mod tidy
RUN go build -o main ./cmd/main.go

CMD ["./main"]
