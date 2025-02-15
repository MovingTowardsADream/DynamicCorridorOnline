package main

import (
	"context"
	"fmt"

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

	fmt.Println(cfg)

	_ = ctx
}
