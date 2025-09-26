package proxy

import (
	"fmt"
	"testing"
)

func TestClose_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &httpServer{
		httpPort: 8080,
	}

	err := server.Close()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
