package httpserver

import (
	"fmt"
	"testing"
)

func TestWithBindAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	s := &Server{}
	WithBindAddr("127.0.0.1:8080")(s)
	if s.bindAddr != "127.0.0.1:8080" {
		t.Errorf("WithBindAddr failed")
	}
}
