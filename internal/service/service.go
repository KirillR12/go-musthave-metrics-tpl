package service

type Storage interface {
	UpdateGauge(name string, value float64)
	UpdateCount(name string, value int64)
	GetGauge(name string) (float64, error)
	GetCount(name string) (int64, error)
}

type MetriceService struct {
	storage Storage
}

func NewMetricsService(storage Storage) *MetriceService {
	return &MetriceService{
		storage: storage,
	}
}

func (m *MetriceService) UpdateCount(name string, value int64) {
	m.storage.UpdateCount(name, value)
}

func (m *MetriceService) UpdateGauge(name string, value float64) {
	m.storage.UpdateGauge(name, value)
}

func (m *MetriceService) GetGaugeMetric(name string) (float64, error) {
	return m.storage.GetGauge(name)
}

func (m *MetriceService) GetCountMetric(name string) (int64, error) {
	return m.storage.GetCount(name)
}
