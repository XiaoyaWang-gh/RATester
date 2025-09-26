package httpserver

import (
	"fmt"
	"testing"
)

func TestBindPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		bindPort: 8080,
	}

	expected := 8080
	actual := server.BindPort()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
