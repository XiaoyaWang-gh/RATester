package orm

import (
	"fmt"
	"testing"
)

func TestnewdbBaseOracle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := newdbBaseOracle()
	if db == nil {
		t.Error("newdbBaseOracle() failed")
	}
}
