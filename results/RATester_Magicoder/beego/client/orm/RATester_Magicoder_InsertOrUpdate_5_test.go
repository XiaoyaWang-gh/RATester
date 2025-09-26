package orm

import (
	"fmt"
	"testing"
)

func TestInsertOrUpdate_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := &DoNothingOrm{}
	_, err := orm.InsertOrUpdate(nil)
	if err != nil {
		t.Errorf("InsertOrUpdate failed, expected nil, got %v", err)
	}
}
