package capture

import (
	"fmt"
	"testing"
)

func TestResponseSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Capture{}
	c.crw = &captureResponseWriter{}
	c.crw.size = 10
	if c.ResponseSize() != 10 {
		t.Errorf("ResponseSize() = %v, want %v", c.ResponseSize(), 10)
	}
}
