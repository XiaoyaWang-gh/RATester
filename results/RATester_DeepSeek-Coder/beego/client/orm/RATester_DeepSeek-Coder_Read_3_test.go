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

	type TestModel struct {
		Id int
	}

	d := &DoNothingOrm{}

	err := d.Read(&TestModel{Id: 1}, "Id")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
