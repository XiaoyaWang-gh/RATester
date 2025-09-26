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
	_, err := orm.Insert(nil)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
