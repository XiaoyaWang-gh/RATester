package client

import (
	"fmt"
	"testing"
)

func TestAddParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &QueryParam{}
	r := map[string][]string{
		"key1": []string{"value1", "value2"},
		"key2": []string{"value3", "value4"},
	}
	p.AddParams(r)
	if p.Len() != 4 {
		t.Errorf("AddParams failed")
	}
}
