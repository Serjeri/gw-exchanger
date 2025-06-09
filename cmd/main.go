package main

import (
	"gw-exchanger/domain/app"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	log.Info("start application")

	application := app.New(50051)

	go func() {
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
