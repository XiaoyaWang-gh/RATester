package mock

import (
	"fmt"
	"testing"
)

func TestOne_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	err := d.One(nil, "")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
