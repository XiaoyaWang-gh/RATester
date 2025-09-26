package orm

import (
	"fmt"
	"testing"
)

func TestnewdbBaseMysql_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := newdbBaseMysql()
	if db == nil {
		t.Error("newdbBaseMysql() failed")
	}
}
