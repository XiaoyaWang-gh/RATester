package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDbTypes_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseSqlite{}
	if got := d.DbTypes(); !reflect.DeepEqual(got, sqliteTypes) {
		t.Errorf("DbTypes() = %v, want %v", got, sqliteTypes)
	}
}
