package metrics

import (
	"fmt"
	"testing"
)

func TestNewPrometheusState_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	ps := newPrometheusState()
	if ps == nil {
		t.Fatal("newPrometheusState() returned nil")
	}
	if ps.dynamicConfig == nil {
		t.Fatal("newPrometheusState() returned a prometheusState with nil dynamicConfig")
	}
	if ps.deletedURLs == nil {
		t.Fatal("newPrometheusState() returned a prometheusState with nil deletedURLs")
	}
}
