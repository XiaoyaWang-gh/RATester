package observability

import (
	"fmt"
	"testing"
)

func TestStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	recorder := &statusCodeRecorder{status: 200}

	if recorder.Status() != 200 {
		t.Errorf("Expected status code to be 200, got %d", recorder.Status())
	}
}
