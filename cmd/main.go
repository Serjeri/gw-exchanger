package main

import (
	"gw-exchanger/domain/repository"
	"gw-exchanger/domain/repository/query"
	"gw-exchanger/domain/services"
	"log"
	"google.golang.org/grpc"
	pb "github.com/Serjeri/proto-exchange/exchange"
)

func main() {
	conn, err := repository.Connect()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	repo := query.NewRepository(conn)
	userService := services.NewUserService(repo)

	grpcServer := grpc.NewServer()
}
