package orm

import (
	"fmt"
	"testing"
)

func TestOne_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var o querySet
	var container interface{}
	var cols []string
	var err error
	err = o.One(container, cols...)
	if err != nil {
		t.Errorf("TestOne error, err:%v", err)
	}
}
