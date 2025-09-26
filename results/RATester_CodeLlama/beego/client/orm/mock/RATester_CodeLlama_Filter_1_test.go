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
	s := "s"
	i := []interface{}{}
	d.Filter(s, i...)
	if d.Filter(s, i...) != d {
		t.Errorf("Filter() = %v, want %v", d.Filter(s, i...), d)
	}
}
