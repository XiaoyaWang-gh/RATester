package client

import (
	"fmt"
	"testing"
)

func TestDelDatas_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &FormData{}
	f.AddDatas(map[string][]string{
		"key1": {"value1"},
		"key2": {"value2"},
		"key3": {"value3"},
	})
	f.DelDatas("key1", "key2")
	if f.Len() != 1 {
		t.Errorf("DelDatas failed")
	}
}
