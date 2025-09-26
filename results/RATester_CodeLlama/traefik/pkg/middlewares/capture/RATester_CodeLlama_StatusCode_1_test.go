package capture

import (
	"fmt"
	"testing"
)

func TestStatusCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Capture{}
	c.crw = &captureResponseWriter{}
	c.crw.status = 200
	if c.StatusCode() != 200 {
		t.Errorf("StatusCode() = %v, want %v", c.StatusCode(), 200)
	}
}
