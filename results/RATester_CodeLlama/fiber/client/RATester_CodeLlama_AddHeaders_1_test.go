package client

import (
	"fmt"
	"testing"
)

func TestAddHeaders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Header{}
	r := map[string][]string{
		"key1": []string{"value1", "value2"},
		"key2": []string{"value3", "value4"},
	}
	h.AddHeaders(r)
	if h.Len() != 2 {
		t.Errorf("AddHeaders failed")
	}
}
