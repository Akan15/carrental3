package main

import (
	"log"
	"net"
	"net/http"

	"github.com/Akan15/carrental3/user-service/internal/handlers"
	"github.com/Akan15/carrental3/user-service/internal/metrics"
	natsPkg "github.com/Akan15/carrental3/user-service/internal/nats"
	"github.com/Akan15/carrental3/user-service/internal/repository"
	"github.com/Akan15/carrental3/user-service/internal/usecase"
	emailPkg "github.com/Akan15/carrental3/user-service/internal/usecase/email"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	pb "github.com/Akan15/carrental3/user-service/proto"

	"github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Загрузка переменных окружения из .env
	_ = godotenv.Load()

	// Подключение к MongoDB
	db := repository.InitMongo()
	repo := repository.NewMongoUserRepo(db)

	// Инициализация NATS Publisher
	natsPkg.InitPublisher()

	// Инициализация usecase и handler
	uc := usecase.NewUserUseCase(repo, emailPkg.SendEmail)
	handler := handlers.NewUserHandler(uc)

	// Инициализация Prometheus метрик
	metrics.Init()

	// Запуск отдельного HTTP-сервера для метрик
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("📊 Метрики Prometheus запущены на :2112/metrics")
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()

	// Запуск gRPC-сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Не удалось запустить gRPC-сервер: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	log.Println("✅ UserService запущен на порту :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Ошибка при запуске gRPC-сервера: %v", err)
	}
}
