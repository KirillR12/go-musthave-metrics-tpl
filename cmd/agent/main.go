package main

import (
	"flag"
	"log"
	"os"
	"strconv"
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

	if addrEnv := os.Getenv("ADDRESS"); addrEnv != "" {
		*addr = addrEnv
	}

	if reportEnv := os.Getenv("REPORT_INTERVAL"); reportEnv != "" {
		value, err := strconv.Atoi(reportEnv)
		if err != nil {
			log.Fatalf("invalid POLL_INTERVAL: %v", err)
		}

		*r = value

	}

	if pollEnv := os.Getenv("POLL_INTERVAL"); pollEnv != "" {
		value, err := strconv.Atoi(pollEnv)
		if err != nil {
			log.Fatalf("invalid POLL_INTERVAL: %v", err)
		}

		*p = value
	}

	reportInt := time.Duration(*r) * time.Second
	pollInt := time.Duration(*p) * time.Second

	a := agent.NewAgent(pollInt, reportInt, client)

	_ = a.Run()
}
