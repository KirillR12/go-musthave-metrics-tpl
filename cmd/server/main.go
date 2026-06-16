package main

import (
	"flag"
	"fmt"

	"github.com/KirillR12/go-musthave-metrics-tpl/internal/app"
	"github.com/KirillR12/go-musthave-metrics-tpl/internal/handler"
	"github.com/KirillR12/go-musthave-metrics-tpl/internal/repository"
	"github.com/KirillR12/go-musthave-metrics-tpl/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	addr := flag.String("a", "localhost:8080", "address to server")

	flag.Parse()

	e := echo.New()
	// cfg := config.NewConfig()

	application := app.NewApp(*addr, e)

	storage := repository.NewMemStorage()
	metriceService := service.NewMetricsService(storage)
	metriceHandler := handler.NewMetricsHandler(metriceService)

	metriceHandler.RegisterRoute(e)

	if err := application.Start(); err != nil {
		fmt.Println("not correct start server")
	}
}
