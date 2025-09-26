package utils

import (
	"fmt"
	"testing"
)

func TestSet_41(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := NewBeeMap()
	k := "key"
	v := "value"
	if !m.Set(k, v) {
		t.Errorf("Set failed")
	}
	if m.Get(k) != v {
		t.Errorf("Set failed")
	}
}
