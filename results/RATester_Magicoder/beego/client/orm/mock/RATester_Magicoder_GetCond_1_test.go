package mock

import (
	"fmt"
	"testing"
)

func TestGetCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	cond := d.GetCond()
	if cond == nil {
		t.Error("Expected non-nil condition, got nil")
	}
}
