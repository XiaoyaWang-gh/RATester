package mock

import (
	"fmt"
	"testing"
)

func TestFilterRaw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	qs := d.FilterRaw("column", "value")
	if _, ok := qs.(*DoNothingQuerySetter); !ok {
		t.Errorf("Expected *DoNothingQuerySetter, got %T", qs)
	}
}
