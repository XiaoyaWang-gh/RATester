package testhelpers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWith_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &CollectingGauge{}
	labelValues := []string{"foo", "bar"}
	g.With(labelValues...)
	if !reflect.DeepEqual(g.LastLabelValues, labelValues) {
		t.Errorf("expected %v, got %v", labelValues, g.LastLabelValues)
	}
}
