package streamserver

import (
	"fmt"
	"testing"
)

func TestBindPort_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{
		bindPort: 1234,
	}
	if s.BindPort() != 1234 {
		t.Fatal("bindPort not match")
	}
}
