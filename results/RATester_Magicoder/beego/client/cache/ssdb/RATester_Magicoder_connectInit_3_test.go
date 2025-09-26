package ssdb

import (
	"fmt"
	"testing"
)

func TestconnectInit_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		conninfo: []string{"localhost:8888"},
	}
	err := cache.connectInit()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
