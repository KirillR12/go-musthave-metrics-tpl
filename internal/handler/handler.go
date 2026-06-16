package handler

import (
	"fmt"
	"net/http"
	"strconv"

	models "github.com/KirillR12/go-musthave-metrics-tpl/internal/model"
	"github.com/labstack/echo/v4"
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

func (h *MetricsHandler) RegisterRoute(e *echo.Echo) {
	e.POST("/update/:type/:name/:value", h.UpdateMetric)
}

func (h *MetricsHandler) UpdateMetric(c echo.Context) error {
	metricType := c.Param("type")
	metricName := c.Param("name")
	metricValue := c.Param("value")

	fmt.Println("received:", metricType, metricName, metricValue)

	if metricName == "" {
		return c.NoContent(http.StatusNotFound)
	}

	switch metricType {
	case models.Counter:
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		h.service.UpdateCount(metricName, value)
		return c.NoContent(http.StatusOK)

	case models.Gauge:
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		h.service.UpdateGauge(metricName, value)
		return c.NoContent(http.StatusOK)

	default:
		return c.NoContent(http.StatusBadRequest)
	}
}
