package app

import (
	"log/slog"
	"sso/internal/app/grpcapp"
	"time"

)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	//TODO: инициализировать storage
	//TODO: init auth service(auth) сервисный слой инициализация
	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCServer: grpcApp,
	}
}
