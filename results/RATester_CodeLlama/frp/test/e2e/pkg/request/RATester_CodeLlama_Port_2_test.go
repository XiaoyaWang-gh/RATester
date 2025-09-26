package request

import (
	"fmt"
	"testing"
)

func TestPort_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.Port(1234)
	if r.port != 1234 {
		t.Errorf("Port() failed")
	}
}
