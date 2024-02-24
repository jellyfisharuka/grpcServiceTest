package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	//TODO: инициализировать конфиг
	cfg := config.MustLoad()
	fmt.Print(cfg)
	//TODO: инициализировать логгер
	//TODO: инициализировать приложение app
	//TODO: запустить gRPC-сервер приложения
}
