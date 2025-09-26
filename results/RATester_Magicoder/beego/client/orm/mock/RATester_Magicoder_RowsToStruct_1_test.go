package mock

import (
	"fmt"
	"testing"
)

func TestRowsToStruct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingRawSetter{}
	_, err := d.RowsToStruct(nil, "", "")
	if err != nil {
		t.Errorf("RowsToStruct returned an error: %v", err)
	}
}
