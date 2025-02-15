package app

import (
	"context"

	userusecase "TicTacToe/internal/application/usecase/user"
	"TicTacToe/internal/infrastructure/config"
	httpserver "TicTacToe/internal/infrastructure/controller/http"
	"TicTacToe/internal/infrastructure/repository/postgres"
	userrepo "TicTacToe/internal/infrastructure/repository/postgres/user"
	"TicTacToe/pkg/hasher"
	"TicTacToe/pkg/logger"
)

type App struct {
	HTTPServer *httpserver.Server
	Storage    *postgres.Postgres
}

func New(ctx context.Context, log logger.Logger, cfg *config.Config) *App {
	storage, err := postgres.New(
		ctx,
		cfg.Storage.URL,
		postgres.MaxPoolSize(cfg.Storage.PoolMax),
		postgres.ConnAttempts(cfg.Storage.ConnAttempts),
		postgres.ConnTimeout(cfg.Storage.ConnTimeout),
	)

	if err != nil {
		panic("storage connection error: " + err.Error())
	}

	err = storage.Ping(ctx)

	if err != nil {
		panic("storage ping error: " + err.Error())
	}

	userRepo := userrepo.New(storage)
	hash := hasher.NewSHA1Hash(cfg.Security.PasswordSalt)
	userUseCase := userusecase.New(log, hash, userRepo, cfg.TokenTLL, cfg.Security.SigningKey)

	return &App{}
}

func (a *App) Run() {

}
