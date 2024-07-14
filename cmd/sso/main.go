package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: инициализация логгера

	// TODO: инициализация приложения (app)

	// TODO: запуск grpc-сервер
}
