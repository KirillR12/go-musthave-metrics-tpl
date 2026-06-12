package service

type Storage interface {
	UpdateGauge(name string, value float64)
	UpdateCount(name string, value int64)
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
