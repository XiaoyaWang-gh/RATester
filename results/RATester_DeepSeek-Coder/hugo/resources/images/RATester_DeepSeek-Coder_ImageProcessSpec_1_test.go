package images

import (
	"fmt"
	"testing"
)

func TestImageProcessSpec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filter := processFilter{
		spec: "test_spec",
	}

	spec := filter.ImageProcessSpec()

	if spec != "test_spec" {
		t.Errorf("Expected 'test_spec', got '%s'", spec)
	}
}
