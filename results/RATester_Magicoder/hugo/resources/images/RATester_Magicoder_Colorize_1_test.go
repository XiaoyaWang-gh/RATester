package images

import (
	"fmt"
	"testing"
)

func TestColorize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}

	hue := 0.5
	saturation := 0.5
	percentage := 0.5

	filter := filters.Colorize(hue, saturation, percentage)

	if filter == nil {
		t.Error("Expected filter to not be nil")
	}
}
