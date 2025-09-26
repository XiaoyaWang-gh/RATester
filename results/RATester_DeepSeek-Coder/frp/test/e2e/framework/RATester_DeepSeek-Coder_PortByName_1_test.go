package framework

import (
	"fmt"
	"testing"
)

func TestPortByName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &Framework{
		usedPorts: map[string]int{
			"test": 8080,
		},
	}

	port := f.PortByName("test")
	if port != 8080 {
		t.Errorf("Expected port 8080, got %d", port)
	}
}
