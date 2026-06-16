package repository

import (
	"fmt"
	"sync"
)

type MemStorage struct {
	mu       sync.RWMutex
	counters map[string]int64
	gauges   map[string]float64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		counters: make(map[string]int64),
		gauges:   make(map[string]float64),
	}
}

func (m *MemStorage) UpdateCount(name string, value int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.counters[name] += value
}

func (m *MemStorage) UpdateGauge(name string, value float64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.gauges[name] = value
}

func (m *MemStorage) GetGauge(name string) (float64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.gauges[name]

	if !ok {
		return 0, fmt.Errorf("not found value gauge: %s", name)
	}

	return value, nil

}

func (m *MemStorage) GetCount(name string) (int64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.counters[name]
	if !ok {
		return 0, fmt.Errorf("not found value count: %s", name)
	}

	return value, nil
}
