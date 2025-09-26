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
	pp.SetParams(map[string]string{
		"key1": "val1",
		"key2": "val2",
	})

	if pp["key1"] != "val1" || pp["key2"] != "val2" {
		t.Errorf("SetParams failed")
	}
}
