package orm

import (
	"fmt"
	"testing"
)

func TestShowColumnsQuery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseOracle{}
	table := "test"
	want := "SELECT COLUMN_NAME FROM ALL_TAB_COLUMNS WHERE TABLE_NAME ='TEST'"
	got := d.ShowColumnsQuery(table)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
