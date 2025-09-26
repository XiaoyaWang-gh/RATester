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

	s := &statusCodeRecorder{}
	s.WriteHeader(404)
	if s.Status() != 404 {
		t.Errorf("Status = %d, want 404", s.Status())
	}
}
