package prometheus

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestOpenConnection_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registry := prometheus.NewRegistry()
	metrics := &serverMetrics{
		connectionCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "connection_count",
			Help: "Number of connections",
		}, []string{"name", "proxyType"}),
	}
	registry.MustRegister(metrics.connectionCount)

	metrics.OpenConnection("test", "test")

	metricFamilies, err := registry.Gather()
	if err != nil {
		t.Fatal(err)
	}

	for _, mf := range metricFamilies {
		if *mf.Name == "connection_count" {
			for _, m := range mf.Metric {
				if *m.Label[0].Value == "test" && *m.Label[1].Value == "test" {
					if *m.Gauge.Value != 1 {
						t.Errorf("Expected connection count to be 1, but got %f", *m.Gauge.Value)
					}
					return
				}
			}
		}
	}
	t.Error("Metric not found")
}
