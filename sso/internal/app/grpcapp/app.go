package grpcapp

import (
	"fmt"
	"log/slog"
	"net"
	authgrpc "sso/internal/grpc/auth"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	Port       int
}

func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer)
	return &App {
		log: log,
		gRPCServer: gRPCServer,
		Port: port,
	} //подключаем обработчик
}
//MustRun runs gRPC server and panics if any error occurs
func (a *App) MustRun() {
	if err:=a.Run(); err!=nil {
		panic("err")
	}
}
func (a *App) Run() error {
	const operation="grpcapp.Run"
	log := a.log.With (
		slog.String("op", operation),
		slog.Int("port", a.Port),
	)
	 l, err:= net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	 if err!=nil {
		return fmt.Errorf("%s:%w", operation, err)
	}
	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))
	if err:=a.gRPCServer.Serve(l); err!=nil {
	return fmt.Errorf("%s: %w", operation, err)
	}
	return nil
}

//Stop stops gRPC server(for graceful shutdown)
func (a *App) Stop() {
	const operation="grpcapp.Stop"
	log := a.log.With(slog.String("op", operation))
	log.Info("stopping gRCPC server", slog.Int("port", a.Port))
	a.gRPCServer.GracefulStop() //прекращает прием новых и ждет пока будут обработаны старые запросы
}