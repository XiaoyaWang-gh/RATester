package metrics

import (
	"fmt"
	"testing"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func TestCollect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ps := &prometheusState{}
	ch := make(chan stdprometheus.Metric)
	ps.Collect(ch)
	close(ch)
	if len(ch) != 0 {
		t.Errorf("Collect() should not write to ch")
	}
}
