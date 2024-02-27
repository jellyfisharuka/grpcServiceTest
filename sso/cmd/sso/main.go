package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"syscall"
)

const (
	envLocal = "locall"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TODO: инициализировать конфиг
	cfg := config.MustLoad()
	fmt.Print(cfg)
	log := setupLogger(cfg.Env)
	log.Info("starting application",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
		slog.Int("port", cfg.GRPC.Port),
	)
	log.Debug("debug message")
	log.Error("error message")
	log.Warn("warn message")

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	go application.GRPCServer.MustRun() //go - запускаем в асинх виде. то есть внутри отдельной горутины
	//TODO: инициализировать логгер
	//TODO: инициализировать приложение app
	//TODO: запустить gRPC-сервер приложения
	//Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT) //signals from operation system ждет эти сигналы и после запишет в канал стоп
	signal :=<-stop //reading from channel - blocking info
	log.Info("stopping application", slog.String("signal", signal.String()))
	//если в канале пусто и мы хотим прочитать из канала то мы зависнем пока канал не будет хоть чем то заполнен 
	application.GRPCServer.Stop()
	log.Info("application stopped")
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
