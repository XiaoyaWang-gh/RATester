package prometheus

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
			Help: "Traffic in",
		}, []string{"name", "proxyType"}),
	}
	registry.MustRegister(metrics.trafficIn)

	testServer := httptest.NewServer(promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	defer testServer.Close()

	metrics.AddTrafficIn("test", "test", 100)

	resp, err := http.Get(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expected := `
	# HELP traffic_in Traffic in
	# TYPE traffic_in counter
	traffic_in{name="test",proxyType="test"} 100
	`
	if string(body) != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, string(body))
	}
}
