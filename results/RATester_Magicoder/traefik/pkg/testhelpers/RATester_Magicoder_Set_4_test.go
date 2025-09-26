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
	g.Set(123.456)

	if g.GaugeValue != 123.456 {
		t.Errorf("Expected GaugeValue to be 123.456, but got %f", g.GaugeValue)
	}
}
