package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"TicTacToe/internal/application/app"
	"TicTacToe/internal/infrastructure/config"
	"TicTacToe/pkg/logger"
)

func main() {
	ctx := context.Background()

	cfg := config.MustLoad()

	log, err := logger.Setup(cfg.Log.Level, cfg.Log.Path)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := log.Close(); err != nil {
			log.Error("failed to close logger", log.Err(err))
		}
	}()

	application := app.New(ctx, log, cfg)

	go func() {
		if errServ := application.HTTPServer.Run(); errServ != nil {
			log.Error("server was shut down due to an error: ", log.Err(errServ))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-stop:
	}

	log.Info("starting graceful shutdown")

	if err := application.HTTPServer.Shutdown(); err != nil {
		log.Error("server was shut down with error: ", log.Err(err))
	}

	application.Storage.Close()

	log.Info("gracefully stopped")
}
