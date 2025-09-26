package orm

import (
	"fmt"
	"testing"
)

func TestValuesList_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var results *[]ParamsList
	var exprs []string
	var err error
	var num int64
	num, err = o.ValuesList(results, exprs...)
	if err != nil {
		t.Errorf("TestValuesList error:%v", err)
	}
	if num != 0 {
		t.Errorf("TestValuesList error:num is not 0")
	}
}
