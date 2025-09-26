package orm

import (
	"fmt"
	"testing"
)

func TestRaw_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := &DoNothingOrm{}
	query := "SELECT * FROM test"
	args := []interface{}{}
	rawSeter := orm.Raw(query, args...)
	if rawSeter != nil {
		t.Errorf("Expected nil, got %v", rawSeter)
	}
}
