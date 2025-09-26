package orm

import (
	"fmt"
	"testing"
)

func TestDelete_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingOrm{}

	type TestStruct struct {
		ID int64
	}

	testStruct := &TestStruct{ID: 1}

	count, err := d.Delete(testStruct)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if count != 0 {
		t.Errorf("Expected count to be 0, got %v", count)
	}
}
