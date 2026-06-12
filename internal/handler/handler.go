package handler

import (
	"net/http"
	"strconv"
	"strings"

	models "github.com/KirillR12/go-musthave-metrics-tpl/internal/model"
)

type MetricsService interface {
	UpdateGauge(name string, value float64)
	UpdateCount(name string, value int64)
}

type MetricsHandler struct {
	service MetricsService
}

func NewMetricsHandler(service MetricsService) *MetricsHandler {
	return &MetricsHandler{
		service: service,
	}
}

func (h *MetricsHandler) RegisterRoute(mux *http.ServeMux) {
	mux.HandleFunc("/update/", h.UpdateMetric)
}

func (h *MetricsHandler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 5 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	metricType := parts[2]
	metricName := parts[3]
	metricValue := parts[4]

	if metricName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch metricType {
	case models.Counter:
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		h.service.UpdateCount(metricName, value)
		w.WriteHeader(http.StatusOK)

	case models.Gauge:
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		h.service.UpdateGauge(metricName, value)
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
