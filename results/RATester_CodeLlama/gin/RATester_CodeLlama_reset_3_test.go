package gin

import (
	"fmt"
	"testing"
)

func TestReset_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &responseWriter{}
	w.reset(nil)
	if w.ResponseWriter != nil {
		t.Errorf("ResponseWriter should be nil")
	}
	if w.size != noWritten {
		t.Errorf("size should be %d", noWritten)
	}
	if w.status != defaultStatus {
		t.Errorf("status should be %d", defaultStatus)
	}
}
