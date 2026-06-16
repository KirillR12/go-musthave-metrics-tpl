package sender

import (
	"fmt"
	"net/http"

	"github.com/ghosind/go-request"
)

type Sender struct {
	client *request.Client
}

func NewSender(client *request.Client) *Sender {
	return &Sender{
		client: client,
	}
}

func (s *Sender) SendGauge(name string, value float64) error {
	url := fmt.Sprintf("/update/gauge/%s/%g", name, value)

	resp, err := s.client.Request(url, request.RequestOptions{
		Headers: map[string][]string{
			"Content-Type": {"text/plain"},
		},
		Method: http.MethodPost,
	})

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
	url := fmt.Sprintf("/update/counter/%s/%d", name, value)

	resp, err := s.client.Request(url, request.RequestOptions{
		Headers: map[string][]string{
			"Content-Type": {"text/plain"},
		},
		Method: http.MethodPost,
	})

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
