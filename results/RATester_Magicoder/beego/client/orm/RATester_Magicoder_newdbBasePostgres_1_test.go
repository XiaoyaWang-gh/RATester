package orm

import (
	"fmt"
	"testing"
)

func TestnewdbBasePostgres_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := newdbBasePostgres()
	if db == nil {
		t.Error("newdbBasePostgres() failed")
	}
}
