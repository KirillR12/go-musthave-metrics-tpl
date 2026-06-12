package app

import (
	"fmt"
	"net/http"
)

type App struct {
	addr string
	mux  *http.ServeMux
}

func NewApp(url string, mux *http.ServeMux) *App {
	return &App{
		addr: url,
		mux:  mux,
	}
}

func (a *App) Start() error {
	fmt.Println("server start on", a.addr)

	return http.ListenAndServe(a.addr, a.mux)
}
