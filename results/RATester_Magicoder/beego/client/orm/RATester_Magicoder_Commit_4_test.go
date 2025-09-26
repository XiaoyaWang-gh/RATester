package orm

import (
	"fmt"
	"testing"
)

func TestCommit_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := &DoNothingTxOrm{}
	err := orm.Commit()
	if err != nil {
		t.Errorf("Commit() failed. Expected nil, got %v", err)
	}
}
