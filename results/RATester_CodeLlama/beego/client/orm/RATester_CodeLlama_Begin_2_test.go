package orm

import (
	"fmt"
	"testing"
)

func TestBegin_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbQueryLog{}
	tx, err := d.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil {
		t.Fatal("tx is nil")
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
