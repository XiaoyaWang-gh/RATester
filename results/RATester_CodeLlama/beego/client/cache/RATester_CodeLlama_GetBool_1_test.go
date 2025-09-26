package cache

import (
	"fmt"
	"testing"
)

func TestGetBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var v interface{}
	if GetBool(v) != false {
		t.Errorf("GetBool(v) = %v, want %v", GetBool(v), false)
	}
}
