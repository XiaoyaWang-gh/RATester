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

	server := &Server{
		bindPort: 8080,
	}

	port := server.BindPort()

	if port != 8080 {
		t.Errorf("Expected port 8080, but got %d", port)
	}
}
