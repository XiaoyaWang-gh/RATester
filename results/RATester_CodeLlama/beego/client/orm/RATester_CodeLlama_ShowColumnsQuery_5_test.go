package orm

import (
	"fmt"
	"testing"
)

func TestShowColumnsQuery_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseSqlite{}
	table := "table"
	want := "pragma table_info('table')"
	got := d.ShowColumnsQuery(table)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
