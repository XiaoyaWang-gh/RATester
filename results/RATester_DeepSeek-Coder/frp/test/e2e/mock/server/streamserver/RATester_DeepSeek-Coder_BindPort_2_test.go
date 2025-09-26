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
		bindPort: 8080,
	}

	port := s.BindPort()
	if port != 8080 {
		t.Errorf("Expected port 8080, got %d", port)
	}
}
