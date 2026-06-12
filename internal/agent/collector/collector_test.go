package collector

import "testing"

func TestCollectRuntimeMetricsFillsGauges(t *testing.T) {
	storage := NewMetricsStorage()

	storage.CollectRuntimeMetrics()

	gauges, _ := storage.Snapshot()

	requiredMetrics := []string{
		"Alloc",
		"BuckHashSys",
		"Frees",
		"GCCPUFraction",
		"GCSys",
		"HeapAlloc",
		"HeapIdle",
		"HeapInuse",
		"HeapObjects",
		"HeapReleased",
		"HeapSys",
		"LastGC",
		"Lookups",
		"MCacheInuse",
		"MCacheSys",
		"MSpanInuse",
		"MSpanSys",
		"Mallocs",
		"NextGC",
		"NumForcedGC",
		"NumGC",
		"OtherSys",
		"PauseTotalNs",
		"StackInuse",
		"StackSys",
		"Sys",
		"TotalAlloc",
		"RandomValue",
	}

	for _, name := range requiredMetrics {
		if _, ok := gauges[name]; !ok {
			t.Fatalf("expected gauge metric %s", name)
		}
	}
}

func TestCollectRuntimeMetricsIncrementsPollCount(t *testing.T) {
	storage := NewMetricsStorage()

	storage.CollectRuntimeMetrics()
	storage.CollectRuntimeMetrics()

	_, counters := storage.Snapshot()

	got := counters["PollCount"]
	want := int64(2)

	if got != want {
		t.Fatalf("expected PollCount %d, got %d", want, got)
	}
}

func TestSnapshotReturnsCopy(t *testing.T) {
	storage := NewMetricsStorage()

	storage.CollectRuntimeMetrics()

	gauges, counters := storage.Snapshot()

	gauges["Alloc"] = -1
	counters["PollCount"] = -1

	newGauges, newCounters := storage.Snapshot()

	if newGauges["Alloc"] == -1 {
		t.Fatal("Snapshot must return copy of gauges map")
	}

	if newCounters["PollCount"] == -1 {
		t.Fatal("Snapshot must return copy of counters map")
	}
}
