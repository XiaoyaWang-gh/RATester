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
	expectedLabelValues := []string{"label1", "label2"}
	g.With(expectedLabelValues...)

	if !reflect.DeepEqual(g.LastLabelValues, expectedLabelValues) {
		t.Errorf("Expected LastLabelValues to be %v, got %v", expectedLabelValues, g.LastLabelValues)
	}
}
