package collector

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"sync"
)

type MetricsStorage struct {
	mu       sync.Mutex
	gauges   map[string]float64
	counters map[string]int64
}

func NewMetricsStorage() *MetricsStorage {
	return &MetricsStorage{
		gauges:   make(map[string]float64),
		counters: make(map[string]int64),
	}
}

func (s *MetricsStorage) CollectRuntimeMetrics() {
	var r runtime.MemStats
	runtime.ReadMemStats(&r)

	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println("collect runtime metrics")

	s.gauges["Alloc"] = float64(r.Alloc)
	s.gauges["BuckHashSys"] = float64(r.BuckHashSys)
	s.gauges["GCCPUFraction"] = r.GCCPUFraction
	s.gauges["GCSys"] = float64(r.GCSys)
	s.gauges["HeapAlloc"] = float64(r.HeapAlloc)
	s.gauges["HeapIdle"] = float64(r.HeapIdle)
	s.gauges["HeapInuse"] = float64(r.HeapInuse)
	s.gauges["HeapObjects"] = float64(r.HeapObjects)
	s.gauges["HeapReleased"] = float64(r.HeapReleased)
	s.gauges["HeapSys"] = float64(r.HeapSys)
	s.gauges["LastGC"] = float64(r.LastGC)
	s.gauges["Lookups"] = float64(r.Lookups)
	s.gauges["MCacheInuse"] = float64(r.MCacheInuse)
	s.gauges["MCacheSys"] = float64(r.MCacheSys)
	s.gauges["Mallocs"] = float64(r.Mallocs)
	s.gauges["NextGC"] = float64(r.NextGC)
	s.gauges["NumForcedGC"] = float64(r.NumForcedGC)
	s.gauges["NumGC"] = float64(r.NumGC)
	s.gauges["OtherSys"] = float64(r.OtherSys)
	s.gauges["PauseTotalNs"] = float64(r.PauseTotalNs)
	s.gauges["StackInuse"] = float64(r.StackInuse)
	s.gauges["StackSys"] = float64(r.StackSys)
	s.gauges["Sys"] = float64(r.Sys)
	s.gauges["TotalAlloc"] = float64(r.TotalAlloc)
	s.gauges["Frees"] = float64(r.Frees)
	s.gauges["MSpanInuse"] = float64(r.MSpanInuse)
	s.gauges["MSpanSys"] = float64(r.MSpanSys)
	s.gauges["RandomValue"] = rand.Float64()

	s.counters["PollCount"]++
}

func (s *MetricsStorage) Snapshot() (map[string]float64, map[string]int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	gauges := make(map[string]float64, len(s.gauges))
	for k, v := range s.gauges {
		gauges[k] = v
	}

	counters := make(map[string]int64, len(s.counters))
	for k, v := range s.counters {
		counters[k] = v
		s.counters[k] = 0
	}

	return gauges, counters
}
