package main

import (
	"gw-exchanger/domain/handlers"
	"gw-exchanger/domain/repository"
	"gw-exchanger/domain/repository/query"
	"gw-exchanger/domain/services"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
)

func main() {
	// cfg := config.MustLoad()

	log.Info("start application")

	conn, err := repository.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	repo := query.NewRepository(conn)
	userService := services.NewUserService(repo)

	grpcServer := grpc.NewServer()
	handlers.Register(grpcServer, userService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Info("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	grpcServer.GracefulStop()
	log.Info("server stopped gracefully")

}
