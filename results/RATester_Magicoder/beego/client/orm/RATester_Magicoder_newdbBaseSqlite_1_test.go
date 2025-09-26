package orm

import (
	"fmt"
	"testing"
)

func TestnewdbBaseSqlite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := newdbBaseSqlite()
	if db == nil {
		t.Error("newdbBaseSqlite() failed")
	}
}
