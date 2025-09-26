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

	db := &dbBaseSqlite{}
	expected := sqliteTypes
	actual := db.DbTypes()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
