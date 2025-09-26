package images

import (
	"fmt"
	"testing"
)

func TestSigmoid_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}

	midpoint := 0.5
	factor := 1.0

	filter := filters.Sigmoid(midpoint, factor)

	if filter == nil {
		t.Error("Expected filter to not be nil")
	}
}
