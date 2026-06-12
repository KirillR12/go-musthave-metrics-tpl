package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockMetricService struct {
	gaugeName    string
	gaugeValue   float64
	counterName  string
	counterValue int64
}

func (m *mockMetricService) UpdateGauge(name string, value float64) {
	m.gaugeName = name
	m.gaugeValue = value
}

func (m *mockMetricService) UpdateCount(name string, value int64) {
	m.counterName = name
	m.counterValue = value
}

func TestUpdateMetricsGaugeSuccess(t *testing.T) {
	service := &mockMetricService{}
	h := NewMetricsHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/gauge/Alloc/123.45", nil)
	w := httptest.NewRecorder()

	h.UpdateMetric(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, res.StatusCode)
	}

	if service.gaugeName != "Alloc" {
		t.Fatalf("expected guage name Alloc, got %s", service.gaugeName)
	}

	if service.gaugeValue != 123.45 {
		t.Fatalf("expected guage value 123.45, got %g", service.gaugeValue)
	}
}

func TestUpdateMetricsCounterSuccess(t *testing.T) {
	service := &mockMetricService{}
	h := NewMetricsHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/counter/PollCount/5", nil)

	w := httptest.NewRecorder()

	h.UpdateMetric(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code: %d, got %d", http.StatusOK, res.StatusCode)
	}

	if service.counterName != "PollCount" {
		t.Errorf("expected counter name PollCount, got %s", service.counterName)
	}

	if service.counterValue != 5 {
		t.Errorf("expected counter value 5, got %d", service.counterValue)
	}
}

func TestUpdateMetricWrongMethod(t *testing.T) {
	service := &mockMetricService{}
	h := NewMetricsHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/update/gauge/Alloc/123.45", nil)
	w := httptest.NewRecorder()

	h.UpdateMetric(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, res.StatusCode)
	}
}

func TestUpdateMetricWithoutName(t *testing.T) {
	service := &mockMetricService{}
	h := NewMetricsHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/gauge//123.45", nil)
	w := httptest.NewRecorder()

	h.UpdateMetric(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, res.StatusCode)
	}
}

func TestUpdateMetricWrongType(t *testing.T) {
	service := &mockMetricService{}
	h := NewMetricsHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/wrong/Alloc/123.45", nil)
	w := httptest.NewRecorder()

	h.UpdateMetric(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestUpdateMetricWrongGaugeValue(t *testing.T) {
	service := &mockMetricService{}
	h := NewMetricsHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/gauge/Alloc/abc", nil)
	w := httptest.NewRecorder()

	h.UpdateMetric(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestUpdateMetricWrongCounterValue(t *testing.T) {
	service := &mockMetricService{}
	h := NewMetricsHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/counter/PollCount/12.5", nil)
	w := httptest.NewRecorder()

	h.UpdateMetric(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}
