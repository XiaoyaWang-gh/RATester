package orm

import (
	"fmt"
	"testing"
)

func TestSupportUpdateJoin_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBasePostgres{}
	if got := d.SupportUpdateJoin(); got != false {
		t.Errorf("dbBasePostgres.SupportUpdateJoin() = %v, want %v", got, false)
	}
}
