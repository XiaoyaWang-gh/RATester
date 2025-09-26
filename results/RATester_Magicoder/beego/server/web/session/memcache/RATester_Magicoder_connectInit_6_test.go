package memcache

import (
	"fmt"
	"testing"
)

func TestconnectInit_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rp := &MemProvider{
		conninfo: []string{"localhost:11211"},
	}

	err := rp.connectInit()
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}
