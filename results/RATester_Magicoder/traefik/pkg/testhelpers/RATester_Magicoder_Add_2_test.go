package testhelpers

import (
	"fmt"
	"testing"
)

func TestAdd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &CollectingGauge{}
	g.Add(1.0)

	if g.GaugeValue != 1.0 {
		t.Errorf("Expected GaugeValue to be 1.0, but got %f", g.GaugeValue)
	}
}
