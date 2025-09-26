package proxy

import (
	"fmt"
	"testing"
)

func TestGetName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &BaseProxy{
		name: "test_name",
	}

	name := pxy.GetName()

	if name != "test_name" {
		t.Errorf("Expected 'test_name', but got '%s'", name)
	}
}
