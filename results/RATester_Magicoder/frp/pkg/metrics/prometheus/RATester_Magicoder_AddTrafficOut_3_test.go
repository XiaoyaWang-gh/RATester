package prometheus

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestAddTrafficOut_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	reg := prometheus.NewRegistry()
	metrics := &serverMetrics{
		trafficOut: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "traffic_out",
		}, []string{"name", "proxyType"}),
	}
	reg.MustRegister(metrics.trafficOut)

	metrics.AddTrafficOut("test", "http", 100)

	metricFam, err := reg.Gather()
	if err != nil {
		t.Fatal(err)
	}

	for _, mf := range metricFam {
		if *mf.Name == "traffic_out" {
			for _, m := range mf.Metric {
				if *m.Label[0].Name == "name" && *m.Label[0].Value == "test" &&
					*m.Label[1].Name == "proxyType" && *m.Label[1].Value == "http" {
					if *m.Counter.Value != 100 {
						t.Errorf("Expected traffic_out to be 100, got %f", *m.Counter.Value)
					}
				}
			}
		}
	}
}
