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

	crw := &captureResponseWriter{}
	crw.size = 100
	if crw.Size() != 100 {
		t.Errorf("Size() = %v, want %v", crw.Size(), 100)
	}
}
