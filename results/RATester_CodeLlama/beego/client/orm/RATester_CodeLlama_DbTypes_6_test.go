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

	d := &dbBasePostgres{}
	if got := d.DbTypes(); !reflect.DeepEqual(got, postgresTypes) {
		t.Errorf("DbTypes() = %v, want %v", got, postgresTypes)
	}
}
