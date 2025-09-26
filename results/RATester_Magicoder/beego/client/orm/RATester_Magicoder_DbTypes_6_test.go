package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDbTypes_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBasePostgres{}
	expected := postgresTypes
	actual := db.DbTypes()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
