package agent

import (
	"time"

	"github.com/KirillR12/go-musthave-metrics-tpl/internal/agent/collector"
	"github.com/KirillR12/go-musthave-metrics-tpl/internal/agent/sender"
	"github.com/ghosind/go-request"
)

type Agent struct {
	storage        *collector.MetricsStorage
	sender         *sender.Sender
	pollInterval   time.Duration
	reportInterval time.Duration
}

func NewAgent(pollInterval time.Duration, reportInterval time.Duration, client *request.Client) *Agent {
	return &Agent{
		storage:        collector.NewMetricsStorage(),
		sender:         sender.NewSender(client),
		pollInterval:   pollInterval,
		reportInterval: reportInterval,
	}
}

func (a *Agent) Run() error {
	go func() {
		for {
			a.storage.CollectRuntimeMetrics()
			time.Sleep(a.pollInterval)
		}
	}()

	for {
		time.Sleep(a.reportInterval)

		gauges, counters := a.storage.Snapshot()
		if err := a.sender.SendAll(gauges, counters); err != nil {
			return err
		}
	}
}
