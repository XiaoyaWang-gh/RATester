package orm

import (
	"fmt"
	"testing"
)

func TestRead_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	orm := &DoNothingOrm{}

	err := orm.Read(nil, "")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
