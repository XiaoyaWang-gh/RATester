package client

import (
	"fmt"
	"testing"
)

func TestDel_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pp := PathParam{"key1": "val1", "key2": "val2"}
	pp.Del("key1")

	if _, ok := pp["key1"]; ok {
		t.Error("Expected key1 to be deleted")
	}
}
