package main

import (
    "log/slog"
    "os"
    "url-sh/internal/config"
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
