package client

import (
	"fmt"
	"testing"
)

func TestSetParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pp := make(PathParam)
	pp.SetParams(map[string]string{"key1": "val1", "key2": "val2"})

	if len(pp) != 2 {
		t.Errorf("Expected length of PathParam to be 2, but got %d", len(pp))
	}

	if pp["key1"] != "val1" {
		t.Errorf("Expected value of key1 to be 'val1', but got '%s'", pp["key1"])
	}

	if pp["key2"] != "val2" {
		t.Errorf("Expected value of key2 to be 'val2', but got '%s'", pp["key2"])
	}
}
