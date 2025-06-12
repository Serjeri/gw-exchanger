package app

import (
	grpcapp "gw-exchanger/domain/app/grpc"
	"gw-exchanger/domain/repository"
	"gw-exchanger/domain/repository/query"
	"gw-exchanger/domain/services"
	"log"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(grpcPort int, db string) *App {
	conn, err := repository.Connect(db)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	repo := query.NewRepository(conn)
	userService := services.NewUserService(repo)

	grpcApp := grpcapp.New(userService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
