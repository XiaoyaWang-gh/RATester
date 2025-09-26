package client

import (
	"fmt"
	"testing"
)

func TestSetDatas_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &FormData{}
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	f.SetDatas(m)
	if f.Len() != 2 {
		t.Errorf("SetDatas failed")
	}
}
