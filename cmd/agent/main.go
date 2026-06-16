package main

import (
	"flag"
	"time"

	"github.com/KirillR12/go-musthave-metrics-tpl/internal/agent"
	"github.com/ghosind/go-request"
)

func main() {
	addr := flag.String("a", "localhost:8080", "address to server")
	r := flag.Int("r", 10, "report interval")
	p := flag.Int("p", 2, "poll interval")

	flag.Parse()

	client := request.New(request.Config{
		BaseURL: "http://" + *addr,
	})

	reportInterval := time.Duration(*r) * time.Second
	pollInterval := time.Duration(*p) * time.Second

	a := agent.NewAgent(pollInterval, reportInterval, client)

	_ = a.Run()
}
