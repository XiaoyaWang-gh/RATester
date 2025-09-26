package gin

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

	w := &responseWriter{
		size:   500,
		status: 200,
	}

	status := w.Status()
	if status != 200 {
		t.Errorf("Expected status to be 200, got %d", status)
	}
}
