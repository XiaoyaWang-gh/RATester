package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDbTypes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseOracle{}
	if got := d.DbTypes(); !reflect.DeepEqual(got, oracleTypes) {
		t.Errorf("DbTypes() = %v, want %v", got, oracleTypes)
	}
}
