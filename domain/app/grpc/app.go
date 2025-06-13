package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"

	authgrpc "gw-exchanger/domain/handlers"
)

type App struct {
	gRPCServer *grpc.Server
	port       string
}

func New(authService authgrpc.UserService, port string) *App{

	grpcServer := grpc.NewServer()

	authgrpc.Register(grpcServer, authService)

	return &App{
		gRPCServer: grpcServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", a.port)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server started ", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	log.Info(slog.String("op", op), "stopping gRPC server", slog.String("port", a.port))

	a.gRPCServer.GracefulStop()
}
