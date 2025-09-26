package mock

import (
	"fmt"
	"testing"
)

func TestQueryRow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingRawSetter{}
	err := d.QueryRow()
	if err != nil {
		t.Errorf("QueryRow() returned an error: %v", err)
	}
}
