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
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
