package sender

import (
	"fmt"
	"net/http"
)

type Sender struct {
	serverAddress string
	client        *http.Client
}

func NewSender(address string) *Sender {
	return &Sender{
		serverAddress: address,
		client:        &http.Client{},
	}
}

func (s *Sender) SendGauge(name string, value float64) error {
	url := fmt.Sprintf("%s/update/gauge/%s/%g", s.serverAddress, name, value)

	req, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/plain")

	resp, err := s.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (s *Sender) SendCounter(name string, value int64) error {
	url := fmt.Sprintf("%s/update/counter/%s/%d", s.serverAddress, name, value)

	req, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/plain")

	resp, err := s.client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (s *Sender) SendAll(gauges map[string]float64, counters map[string]int64) error {
	fmt.Println("send gauges count:", len(gauges))
	fmt.Println("send counters count:", len(counters))

	for name, value := range gauges {
		if err := s.SendGauge(name, value); err != nil {
			return err
		}
	}

	for name, value := range counters {
		if err := s.SendCounter(name, value); err != nil {
			return err
		}
	}

	return nil
}
