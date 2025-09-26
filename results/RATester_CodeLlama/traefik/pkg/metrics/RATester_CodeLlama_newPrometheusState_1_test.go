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

	ps := newPrometheusState()
	if ps == nil {
		t.Errorf("newPrometheusState() returned nil")
	}
	if ps.dynamicConfig == nil {
		t.Errorf("newPrometheusState() returned nil dynamicConfig")
	}
	if ps.deletedURLs == nil {
		t.Errorf("newPrometheusState() returned nil deletedURLs")
	}
}
