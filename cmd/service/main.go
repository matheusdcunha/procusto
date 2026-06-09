package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/matheusdcunha/procusto/internal/api"
)

func main() {

	if err := run(); err != nil {
		slog.Error("fatal", "error", err)
		os.Exit(1)
	}
}

func run() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	st := ""

	e := echo.New()
	e.Logger = logger
	srv := api.NewServer(st, logger)

	srv.RegisterRoutes(e)

	sc := echo.StartConfig{Address: ":8081"}
	return sc.Start(context.Background(), e)
}
