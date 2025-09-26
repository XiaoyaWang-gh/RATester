package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFloat64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var dest float64
	b := new(ValueBinder)
	b.floatValue("sourceParam", &dest, 64, false)

	if !reflect.DeepEqual(dest, float64(0)) {
		t.Errorf("expected %v, got %v", float64(0), dest)
	}
}
