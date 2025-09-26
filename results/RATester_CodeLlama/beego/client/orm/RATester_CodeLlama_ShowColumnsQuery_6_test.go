package orm

import (
	"fmt"
	"testing"
)

func TestShowColumnsQuery_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBase{}
	table := "table"
	want := "show columns from table"
	got := d.ShowColumnsQuery(table)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
