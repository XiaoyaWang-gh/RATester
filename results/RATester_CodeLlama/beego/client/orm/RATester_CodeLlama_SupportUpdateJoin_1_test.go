package orm

import (
	"fmt"
	"testing"
)

func TestSupportUpdateJoin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseSqlite{}
	if got := d.SupportUpdateJoin(); got != false {
		t.Errorf("dbBaseSqlite.SupportUpdateJoin() = %v, want %v", got, false)
	}
}
