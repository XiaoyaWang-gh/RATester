package orm

import (
	"fmt"
	"testing"
)

func TestAll_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var container interface{}
	var cols []string
	var result int64
	var err error
	result, err = o.All(container, cols...)
	if err != nil {
		t.Errorf("TestAll error, err:%v", err)
	}
	if result != 0 {
		t.Errorf("TestAll error, result:%v", result)
	}
}
