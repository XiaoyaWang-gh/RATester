package observability

import (
	"fmt"
	"testing"
)

func TestWriteHeader_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &statusCodeRecorder{}
	s.WriteHeader(200)
	if s.Status() != 200 {
		t.Errorf("Status() = %d, want %d", s.Status(), 200)
	}
}
