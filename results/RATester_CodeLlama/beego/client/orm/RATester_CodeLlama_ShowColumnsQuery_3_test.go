package orm

import (
	"fmt"
	"testing"
)

func TestShowColumnsQuery_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBasePostgres{}
	table := "table"
	want := "SELECT column_name, data_type, is_nullable FROM information_schema.Columns where table_schema NOT IN ('pg_catalog', 'information_schema') and table_name = 'table'"
	got := d.ShowColumnsQuery(table)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
