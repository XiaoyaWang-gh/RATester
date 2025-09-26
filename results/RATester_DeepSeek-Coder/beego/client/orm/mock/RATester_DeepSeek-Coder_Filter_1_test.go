package mock

import (
	"fmt"
	"testing"
)

func TestFilter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	qs := d.Filter("UserName", "slene")
	if qs == nil {
		t.Errorf("Filter() returned nil")
	}
}
