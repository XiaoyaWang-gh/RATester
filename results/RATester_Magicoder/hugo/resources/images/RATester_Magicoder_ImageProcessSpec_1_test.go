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

	filter := processFilter{spec: "test spec"}
	spec := filter.ImageProcessSpec()
	if spec != "test spec" {
		t.Errorf("Expected 'test spec', but got '%s'", spec)
	}
}
