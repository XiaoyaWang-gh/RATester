package dynacache

import (
	"fmt"
	"testing"
)

func TestWeightFraction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := OptionsPartition{
		Weight: 100,
	}
	if got := o.WeightFraction(); got != 1 {
		t.Errorf("WeightFraction() = %v, want %v", got, 1)
	}
}
