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
		t.Error("newPrometheusState() should not return nil")
	}
	if ps.dynamicConfig == nil {
		t.Error("newPrometheusState().dynamicConfig should not be nil")
	}
	if ps.deletedURLs == nil {
		t.Error("newPrometheusState().deletedURLs should not be nil")
	}
}
