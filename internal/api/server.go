package api

import "log/slog"

type Server struct {
	store  any
	logger *slog.Logger
}

func NewServer(s any, l *slog.Logger) *Server {
	return &Server{store: s, logger: l}
}
