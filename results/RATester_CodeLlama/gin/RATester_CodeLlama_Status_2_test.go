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

	w := &responseWriter{status: 200}
	if w.Status() != 200 {
		t.Errorf("w.Status() = %d; want 200", w.Status())
	}
}
