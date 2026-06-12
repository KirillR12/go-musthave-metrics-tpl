package main

import (
	"time"

	"github.com/KirillR12/go-musthave-metrics-tpl/internal/agent"
)

func main() {
	a := agent.NewAgent("http://localhost:8080", 2*time.Second, 10*time.Second)

	_ = a.Run()
}
