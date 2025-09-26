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

	ormer := &DoNothingTxOrm{}
	err := ormer.Rollback()
	if err != nil {
		t.Errorf("Rollback() failed. Expected nil, got %v", err)
	}
}
