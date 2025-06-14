package app

import (
	grpcapp "gw-exchanger/domain/app/grpc"
	"gw-exchanger/domain/repository"
	"gw-exchanger/domain/repository/query"
	"gw-exchanger/domain/services"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort string, db string) *App {
	conn, err := repository.Connect(db)
	if err != nil {
		panic(err)
	}

	repo := query.NewRepository(conn)
	userService := services.NewUserService(log ,repo)

	grpcApp := grpcapp.New(log, userService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
