package prometheus

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestAddTrafficIn_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	registry := prometheus.NewRegistry()
	metrics := &serverMetrics{
		trafficIn: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "traffic_in",
		}, []string{"name", "proxyType"}),
	}
	registry.MustRegister(metrics.trafficIn)

	metrics.AddTrafficIn("test", "http", 100)

	metricFamilies, err := registry.Gather()
	if err != nil {
		t.Fatal(err)
	}

	for _, mf := range metricFamilies {
		if *mf.Name == "traffic_in" {
			for _, m := range mf.Metric {
				if *m.Label[0].Value == "test" && *m.Label[1].Value == "http" {
					if *m.Counter.Value != 100 {
						t.Errorf("Expected traffic in to be 100, but got %f", *m.Counter.Value)
					}
				}
			}
		}
	}
}
