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

	d := &DoNothingTxOrm{}
	err := d.Commit()
	if err != nil {
		t.Errorf("Commit() failed, got %v", err)
	}
}
