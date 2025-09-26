package mock

import (
	"fmt"
	"testing"
)

func TestRowsToStruct_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	_, err := d.RowsToStruct(nil, "", "")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
