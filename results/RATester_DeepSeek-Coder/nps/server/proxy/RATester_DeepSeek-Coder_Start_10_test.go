package proxy

import (
	"fmt"
	"testing"
)

func TestStart_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &httpServer{
		httpPort: 8080,
	}

	err := s.Start()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
