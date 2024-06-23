package main

import (
    "log/slog"
    "os"
    "url-sh/internal/config"
    "url-sh/internal/storage/sqlite"
)

const (
    envLocal = "local"
    envProd  = "prod"
    envDev   = "dev"
)

func main() {
    cfg := config.MustLoad()

    log := setupLogger(cfg.Env)

    log.Info("Starting server", slog.String("env", cfg.Env))
    log.Debug("Debug messages are enabled")

    storage, err := sqlite.New(cfg.StoragePath)

    if err != nil {
        log.Error("Failed to initialise db", err)
        os.Exit(1)
    }

    _ = storage

}

func setupLogger(env string) *slog.Logger {
    var log *slog.Logger
    switch env {
    case envLocal:
        log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
    case envDev:
        log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
    case envProd:
        log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
    }

    return log
}
