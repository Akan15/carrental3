FROM golang:1.23

WORKDIR /app

# Копируем go.mod и go.sum отдельно
COPY go.mod go.sum ./

# Копируем локальные сервисы в рабочую директорию
COPY ../user-service ./user-service
COPY ../car-service ./car-service
COPY ../rental-service ./rental-service

# Копируем остальные файлы самого api-gateway
COPY . .

# Загружаем зависимости и билдим
RUN go mod tidy
RUN go build -o main ./cmd/main.go

CMD ["./main"]
