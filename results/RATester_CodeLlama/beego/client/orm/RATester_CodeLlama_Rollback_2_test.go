package orm

import (
	"fmt"
	"testing"
)

func TestRollback_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingTxOrm{}
	err := d.Rollback()
	if err != nil {
		t.Errorf("Rollback() = %v, want nil", err)
	}
}
