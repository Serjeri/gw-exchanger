package main

import (
	"gw-exchanger/domain/app"
	"gw-exchanger/domain/config"
	"os"
	"os/signal"
	"syscall"

	"log/slog"
)


func main() {
	cfg := config.MustLoad()

	log := setupLogger()

	log.Info("start application")

	application := app.New(log,cfg.Addressgrpc, cfg.Dburl)

	go func() {
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}

func setupLogger() *slog.Logger {
	var log *slog.Logger

	log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),)

	return log
}
