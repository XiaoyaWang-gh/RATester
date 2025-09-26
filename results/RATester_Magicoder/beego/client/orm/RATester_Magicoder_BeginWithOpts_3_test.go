package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestBeginWithOpts_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := &DoNothingOrm{}
	opts := &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}
	tx, err := orm.BeginWithOpts(opts)
	if err != nil {
		t.Errorf("BeginWithOpts returned an error: %v", err)
	}
	if tx == nil {
		t.Error("BeginWithOpts returned nil TxOrmer")
	}
}
