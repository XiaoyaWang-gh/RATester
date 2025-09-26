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
	g.With("label1", "label2")

	if !reflect.DeepEqual(g.LastLabelValues, []string{"label1", "label2"}) {
		t.Errorf("Expected LastLabelValues to be ['label1', 'label2'], but got %v", g.LastLabelValues)
	}
}
