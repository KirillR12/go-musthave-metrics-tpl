package main

import (
	"fmt"
	"net/http"

	"github.com/KirillR12/go-musthave-metrics-tpl/internal/app"
	"github.com/KirillR12/go-musthave-metrics-tpl/internal/config"
	"github.com/KirillR12/go-musthave-metrics-tpl/internal/handler"
	"github.com/KirillR12/go-musthave-metrics-tpl/internal/repository"
	"github.com/KirillR12/go-musthave-metrics-tpl/internal/service"
)

func main() {
	cfg := config.NewConfig()

	mux := http.NewServeMux()
	application := app.NewApp(cfg.Address, mux)

	storage := repository.NewMemStorage()
	metriceService := service.NewMetriceService(storage)
	metriceHandler := handler.NewMetricsHandler(metriceService)

	metriceHandler.RegisterRoute(mux)

	if err := application.Start(); err != nil {
		fmt.Println("not correct start server")
	}
}
