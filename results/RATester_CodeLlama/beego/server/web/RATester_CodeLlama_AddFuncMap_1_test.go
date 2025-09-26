package web

import (
	"fmt"
	"testing"
)

func TestAddFuncMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := "test"
	fn := func() {}
	err := AddFuncMap(key, fn)
	if err != nil {
		t.Errorf("AddFuncMap failed, err: %v", err)
	}
}
