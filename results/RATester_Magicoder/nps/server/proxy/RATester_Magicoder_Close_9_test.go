package proxy

import (
	"fmt"
	"testing"
)

func TestClose_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &WebServer{}
	err := server.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
