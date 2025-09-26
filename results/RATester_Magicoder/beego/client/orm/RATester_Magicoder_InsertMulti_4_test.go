package orm

import (
	"fmt"
	"testing"
)

func TestInsertMulti_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := &DoNothingOrm{}
	bulk := 10
	mds := []interface{}{"test1", "test2", "test3"}
	_, err := orm.InsertMulti(bulk, mds)
	if err != nil {
		t.Errorf("InsertMulti failed, got error: %v", err)
	}
}
