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
	if got := d.GetCond(); got != nil {
		t.Errorf("DoNothingQuerySetter.GetCond() = %v, want nil", got)
	}
}
