package capture

import (
	"fmt"
	"testing"
)

func TestResponseSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	capture := &Capture{
		crw: &captureResponseWriter{
			size: 100,
		},
	}

	size := capture.ResponseSize()

	if size != 100 {
		t.Errorf("Expected size to be 100, but got %d", size)
	}
}
