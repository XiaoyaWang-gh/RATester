package orm

import (
	"fmt"
	"testing"
)

func TestBegin_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := &DoNothingOrm{}
	tx, err := orm.Begin()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if tx != nil {
		t.Errorf("Expected nil transaction, but got %v", tx)
	}
}
