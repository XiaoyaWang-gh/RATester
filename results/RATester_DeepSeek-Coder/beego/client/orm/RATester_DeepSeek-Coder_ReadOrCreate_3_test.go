package orm

import (
	"fmt"
	"testing"
)

func TestReadOrCreate_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingOrm{}

	type TestModel struct {
		Id   int64
		Name string
	}

	testModel := &TestModel{
		Name: "test",
	}

	created, id, err := d.ReadOrCreate(testModel, "Name")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if created && id == 0 {
		t.Errorf("Expected id to be greater than 0")
	}

	if !created && id == 0 {
		t.Errorf("Expected id to be greater than 0")
	}
}
