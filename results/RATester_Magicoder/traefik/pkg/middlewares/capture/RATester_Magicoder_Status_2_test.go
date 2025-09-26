package capture

import (
	"fmt"
	"testing"
)

func TestStatus_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	crw := &captureResponseWriter{status: 200}
	if crw.Status() != 200 {
		t.Errorf("Expected status 200, got %d", crw.Status())
	}
}
