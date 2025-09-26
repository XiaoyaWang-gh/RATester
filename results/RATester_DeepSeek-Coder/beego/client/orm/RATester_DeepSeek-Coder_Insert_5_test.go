package orm

import (
	"fmt"
	"testing"
)

func TestInsert_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := &DoNothingOrm{}

	type TestModel struct {
		Id int
	}

	id, err := orm.Insert(&TestModel{Id: 1})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if id != 0 {
		t.Errorf("Expected id to be 0, got %v", id)
	}
}
