package main

import (
	"log"
	"net"

	"github.com/Akan15/carrental3/user-service/internal/handlers"
	natsPkg "github.com/Akan15/carrental3/user-service/internal/nats"
	"github.com/Akan15/carrental3/user-service/internal/repository"
	"github.com/Akan15/carrental3/user-service/internal/usecase"
	emailPkg "github.com/Akan15/carrental3/user-service/internal/usecase/email"

	pb "github.com/Akan15/carrental3/user-service/proto"

	"github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	_ = godotenv.Load() // ⬅️ это строка загружает SMTP_FROM и SMTP_PASS

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Подключение MongoDB
	db := repository.InitMongo()
	repo := repository.NewMongoUserRepo(db)

	// ⬇️ Инициализация NATS
	natsPkg.InitPublisher()

	uc := usecase.NewUserUseCase(repo, emailPkg.SendEmail)
	handler := handlers.NewUserHandler(uc)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, handler)

	reflection.Register(grpcServer)

	log.Println("✅ UserService is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
