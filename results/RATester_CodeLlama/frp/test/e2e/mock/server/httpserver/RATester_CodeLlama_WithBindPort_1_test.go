package httpserver

import (
	"fmt"
	"testing"
)

func TestWithBindPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var s Server
	WithBindPort(8080)(&s)
	if s.BindPort() != 8080 {
		t.Errorf("WithBindPort(8080)(&s) failed")
	}
}
