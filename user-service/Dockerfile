FROM golang:1.24.1-alpine


WORKDIR /app

# Копируем все исходники, включая handlers/usecase
COPY . .

RUN go mod download
RUN go build -o main ./cmd/main.go

CMD ["./main"]
