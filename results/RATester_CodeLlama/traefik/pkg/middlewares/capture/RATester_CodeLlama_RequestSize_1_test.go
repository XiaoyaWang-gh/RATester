package capture

import (
	"fmt"
	"testing"
)

func TestRequestSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Capture{}
	if c.RequestSize() != 0 {
		t.Errorf("RequestSize() = %v, want %v", c.RequestSize(), 0)
	}
}
