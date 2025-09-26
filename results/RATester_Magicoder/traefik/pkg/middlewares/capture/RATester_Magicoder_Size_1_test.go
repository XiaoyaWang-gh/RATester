package capture

import (
	"fmt"
	"testing"
)

func TestSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	crw := &captureResponseWriter{
		size: 100,
	}

	if crw.Size() != 100 {
		t.Errorf("Expected size to be 100, but got %d", crw.Size())
	}
}
