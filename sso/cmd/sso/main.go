package main

import (
	"log/slog"
	"os"
	"sso/internal/app"
	"sso/internal/config"
)
const (
	envLocal = "local"
	envDev="dev"
	envProd="prod"
)
func main() {
	//TODO: инициализировать конфиг
	cfg := config.MustLoad()
	log:=setupLogger(cfg.Env)
	log.Info("starting application",
	 slog.String("env", cfg.Env),
	 slog.Any("cfg", cfg),
	 slog.Int("port", cfg.Grpc.Port),
	)
	log.Debug("debug message")
	log.Error("error message")
	log.Warn("warn message")
	application:= app.New(log, cfg.Grpc.Port, cfg.StoragePath, cfg.TokenTTL)
	application.GRPCServer.MustRun()
	//TODO: инициализировать логгер
	//TODO: инициализировать приложение app
	//TODO: запустить gRPC-сервер приложения
}
func setupLogger(env string) *slog.Logger {
var log *slog.Logger
switch env {
case envLocal:
	log=slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level:slog.LevelDebug}),
	)
case envDev:
	log=slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
case envProd:
	log=slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level:slog.LevelInfo}),
	)
}
return log
}
