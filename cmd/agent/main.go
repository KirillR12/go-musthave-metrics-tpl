package main

import (
	"time"

	"github.com/KirillR12/go-musthave-metrics-tpl/internal/agent"
	"github.com/ghosind/go-request"
)

func main() {
	client := request.New(request.Config{
		BaseURL: "http://localhost:8080",
	})

	a := agent.NewAgent(2*time.Second, 10*time.Second, client)

	_ = a.Run()
}
