package http

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"TicTacToe/pkg/logger"
	"TicTacToe/pkg/utils"
)

const (
	_defaultReadTimeout     = 10 * time.Second
	_defaultWriteTimeout    = 10 * time.Second
	_defaultAddr            = 8080
	_defaultShutdownTimeout = 5 * time.Second
)

type Server struct {
	httpServer *http.Server
	log        *slog.Logger

	port            int
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
}

func New(log *slog.Logger, handler http.Handler, opts ...Option) *Server {
	s := &Server{
		log:             log,
		port:            _defaultAddr,
		readTimeout:     _defaultReadTimeout,
		writeTimeout:    _defaultWriteTimeout,
		shutdownTimeout: _defaultShutdownTimeout,
	}

	s.httpServer = &http.Server{
		Handler:      handler,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
		Addr:         utils.FormatAddress("", s.port),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic("cannot run http server: " + err.Error())
	}
}

func (s *Server) Run() error {
	const op = "http - Server.Run"

	l, err := net.Listen("tcp", s.httpServer.Addr)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	s.log.Info("gRPC server started", logger.AnyAttr("addr", l.Addr().String()))

	if err := s.httpServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	const op = "httpServer.Shutdown"

	s.log.With(slog.String("op", op)).
		Info("stopping http server", slog.String("port", s.httpServer.Addr))

	return fmt.Errorf("%s: %w", op, s.httpServer.Shutdown(ctx))
}
