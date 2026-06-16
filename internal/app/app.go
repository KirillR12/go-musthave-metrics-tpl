package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type App struct {
	addr string
	e    *echo.Echo
}

func NewApp(url string, e *echo.Echo) *App {
	return &App{
		addr: url,
		e:    e,
	}
}

func (a *App) Start() error {
	fmt.Println("server start on", a.addr)

	if err := a.e.Start(a.addr); err != nil {
		return fmt.Errorf("error start server")
	}

	return nil
}
