package images

import (
	"fmt"
	"testing"
)

func TestGamma_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}
	gamma := 2.0
	filter := filters.Gamma(gamma)

	if filter == nil {
		t.Error("Expected a filter, but got nil")
	}
}
