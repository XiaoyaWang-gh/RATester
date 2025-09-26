package orm

import (
	"fmt"
	"testing"
)

func TestShowColumnsQuery_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseTidb{}
	table := "table"
	want := "SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE FROM information_schema.Columns WHERE table_schema = DATABASE() AND table_name = 'table'"
	got := d.ShowColumnsQuery(table)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
