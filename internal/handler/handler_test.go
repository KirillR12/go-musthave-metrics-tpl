package handler

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
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

func (m *mockMetricService) GetGaugeMetric(name string) (float64, error) {
	if name != m.gaugeName {
		return 0, errors.New("gauge metric not found")
	}

	return m.gaugeValue, nil
}

func (m *mockMetricService) GetCountMetric(name string) (int64, error) {
	if name != m.counterName {
		return 0, errors.New("counter metric not found")
	}

	return m.counterValue, nil
}

func setupTestHandler(service *mockMetricService) *echo.Echo {
	e := echo.New()

	h := NewMetricsHandler(service)
	h.RegisterRoute(e)

	return e
}

func TestUpdateMetricsGaugeSuccess(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/gauge/Alloc/123.45", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, res.StatusCode)
	}

	if service.gaugeName != "Alloc" {
		t.Fatalf("expected gauge name Alloc, got %s", service.gaugeName)
	}

	if service.gaugeValue != 123.45 {
		t.Fatalf("expected gauge value 123.45, got %g", service.gaugeValue)
	}
}

func TestUpdateMetricsCounterSuccess(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/counter/PollCount/5", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, res.StatusCode)
	}

	if service.counterName != "PollCount" {
		t.Fatalf("expected counter name PollCount, got %s", service.counterName)
	}

	if service.counterValue != 5 {
		t.Fatalf("expected counter value 5, got %d", service.counterValue)
	}
}

func TestGetGaugeMetricSuccess(t *testing.T) {
	service := &mockMetricService{
		gaugeName:  "Alloc",
		gaugeValue: 123.45,
	}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/value/gauge/Alloc", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, res.StatusCode)
	}

	if string(body) != "123.45" {
		t.Fatalf("expected body 123.45, got %s", string(body))
	}
}

func TestGetCounterMetricSuccess(t *testing.T) {
	service := &mockMetricService{
		counterName:  "PollCount",
		counterValue: 5,
	}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/value/counter/PollCount", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, res.StatusCode)
	}

	if string(body) != "5" {
		t.Fatalf("expected body 5, got %s", string(body))
	}
}

func TestGetMetricUnknownGauge(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/value/gauge/Unknown", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, res.StatusCode)
	}
}

func TestGetMetricUnknownCounter(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/value/counter/Unknown", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, res.StatusCode)
	}
}

func TestGetMetricWrongType(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/value/wrong/Alloc", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestUpdateMetricWrongMethod(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/update/gauge/Alloc/123.45", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, res.StatusCode)
	}
}

func TestUpdateMetricWithoutName(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/gauge//123.45", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, res.StatusCode)
	}
}

func TestUpdateMetricWrongType(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/wrong/Alloc/123.45", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestUpdateMetricWrongGaugeValue(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/gauge/Alloc/abc", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestUpdateMetricWrongCounterValue(t *testing.T) {
	service := &mockMetricService{}
	e := setupTestHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/update/counter/PollCount/12.5", nil)
	w := httptest.NewRecorder()

	e.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}
