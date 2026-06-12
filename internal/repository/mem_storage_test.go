package repository

import "testing"

func TestUpdateGaugeReplacesValue(t *testing.T) {
	storage := NewMemStorage()

	storage.UpdateGauge("Alloc", 100.5)
	storage.UpdateGauge("Alloc", 200.5)

	got := storage.gauges["Alloc"]
	want := 200.5

	if got != want {
		t.Fatalf("expected %f, got %f", want, got)
	}
}

func TestUpdateCounterAddsValue(t *testing.T) {
	storage := NewMemStorage()

	storage.UpdateCount("PollCount", 5)
	storage.UpdateCount("PollCount", 3)

	got := storage.counters["PollCount"]
	want := int64(8)

	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}
