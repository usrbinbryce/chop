package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/usrbinbryce/chop/internal/env"
)

func main() {
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config{
		addr: env.GetString("PORT", ":8080"),
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=root password=password dbname=chop sslmode=disable"),
		},
	}

	conn, err := pgxpool.New(ctx, cfg.db.dsn)
	if err != nil {
		slog.Error("failed to connect to db", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	if err := conn.Ping(ctx); err != nil {
		slog.Error("failed to ping db", "error", nil)
	}
	logger.Info("connected to db successfully")

	app := application{
		config: cfg,
		db:     conn,
	}
	if err := app.run(app.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
