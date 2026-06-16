package sender

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ghosind/go-request"
)

func setupNewClient(url string) *request.Client {
	return request.New(request.Config{
		BaseURL: url,
	})
}

func TestSendGauge(t *testing.T) {
	var gotMethod string
	var gotPath string
	var gotContentType string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotContentType = r.Header.Get("Content-Type")

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := setupNewClient(server.URL)
	sender := NewSender(c)

	err := sender.SendGauge("Alloc", 123.45)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Fatalf("expected method POST, got %s", gotMethod)
	}

	if gotPath != "/update/gauge/Alloc/123.45" {
		t.Fatalf("expected path /update/gauge/Alloc/123.45, got %s", gotPath)
	}

	if gotContentType != "text/plain" {
		t.Fatalf("expected Content-Type text/plain, got %s", gotContentType)
	}
}

func TestSendCounter(t *testing.T) {
	var gotMethod string
	var gotPath string
	var gotContentType string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotContentType = r.Header.Get("Content-Type")

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := setupNewClient(server.URL)
	sender := NewSender(c)

	err := sender.SendCounter("PollCount", 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Fatalf("expected method POST, got %s", gotMethod)
	}

	if gotPath != "/update/counter/PollCount/5" {
		t.Fatalf("expected path /update/counter/PollCount/5, got %s", gotPath)
	}

	if gotContentType != "text/plain" {
		t.Fatalf("expected Content-Type text/plain, got %s", gotContentType)
	}
}

func TestSendGaugeReturnsErrorOnBadStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	c := setupNewClient(server.URL)

	sender := NewSender(c)

	err := sender.SendGauge("Alloc", 123.45)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
