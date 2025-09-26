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

	f := processFilter{
		spec: "test",
	}
	if f.ImageProcessSpec() != "test" {
		t.Errorf("ImageProcessSpec() = %v, want %v", f.ImageProcessSpec(), "test")
	}
}
