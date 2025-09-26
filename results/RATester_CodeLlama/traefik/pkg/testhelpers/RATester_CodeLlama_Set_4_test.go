package testhelpers

import (
	"fmt"
	"testing"
)

func TestSet_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &CollectingGauge{}
	value := 1.0
	g.Set(value)
	if g.GaugeValue != value {
		t.Errorf("GaugeValue = %v, want %v", g.GaugeValue, value)
	}
}
