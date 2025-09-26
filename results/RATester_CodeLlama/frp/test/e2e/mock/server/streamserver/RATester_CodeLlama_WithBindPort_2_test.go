package streamserver

import (
	"fmt"
	"testing"
)

func TestWithBindPort_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var s Server
	WithBindPort(1234)(&s)
	if s.BindPort() != 1234 {
		t.Errorf("WithBindPort(1234)(&s) failed")
	}
}
