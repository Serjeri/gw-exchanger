package main

import (
	"gw-exchanger/domain/app"
	"gw-exchanger/domain/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	cfg := config.MustLoad()

	log.Info("start application")

	application := app.New(cfg.Addressgrpc, cfg.Dburl)

	go func() {
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
