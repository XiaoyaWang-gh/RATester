package mock

import (
	"fmt"
	"testing"
)

func TestDistinct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	d := &DoNothingQuerySetter{}
	qs := d.Distinct()

	if _, ok := qs.(*DoNothingQuerySetter); !ok {
		t.Errorf("Expected *DoNothingQuerySetter, got %T", qs)
	}
}
