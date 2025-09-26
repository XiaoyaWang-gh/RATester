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

	p := PathParam{}
	m := map[string]string{"key1": "val1", "key2": "val2"}
	p.SetParams(m)
	if len(p) != 2 {
		t.Errorf("SetParams failed")
	}
}
