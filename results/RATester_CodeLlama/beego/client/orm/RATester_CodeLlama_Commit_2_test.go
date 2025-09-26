package orm

import (
	"fmt"
	"testing"
)

func TestCommit_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	db := &TxDB{}
	err := db.Commit()
	if err != nil {
		t.Errorf("Commit() failed: %v", err)
	}
}
