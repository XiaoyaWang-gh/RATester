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

	testServer := httptest.NewServer(promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	defer testServer.Close()

	metrics.OpenConnection("test", "test")

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
	# HELP connection_count Number of connections
	# TYPE connection_count gauge
	connection_count{name="test",proxyType="test"} 1
	`
	if string(body) != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, string(body))
	}
}
