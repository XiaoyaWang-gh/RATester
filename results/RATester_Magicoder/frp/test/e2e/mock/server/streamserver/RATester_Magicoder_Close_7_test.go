package streamserver

import (
	"fmt"
	"testing"
)

func TestClose_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		netType:  TCP,
		bindAddr: "localhost",
		bindPort: 8080,
	}

	err := server.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
