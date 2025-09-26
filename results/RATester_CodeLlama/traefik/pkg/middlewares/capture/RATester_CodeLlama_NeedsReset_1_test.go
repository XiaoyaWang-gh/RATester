package capture

import (
	"fmt"
	"testing"
)

func TestNeedsReset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Capture{}
	rw := &captureResponseWriter{}
	if c.NeedsReset(rw) {
		t.Error("NeedsReset() should return false")
	}
}
