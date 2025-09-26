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

	ws := &WebServer{}
	err := ws.Close()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
