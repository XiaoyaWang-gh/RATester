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
	err := d.QueryRow(nil)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
