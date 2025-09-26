package orm

import (
	"fmt"
	"testing"
)

func TestBegin_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DB{}
	tx, err := d.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if tx == nil {
		t.Fatal("expected tx")
	}
}
