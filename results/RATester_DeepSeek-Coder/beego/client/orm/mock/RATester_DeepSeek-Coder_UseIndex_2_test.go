package mock

import (
	"fmt"
	"testing"
)

func TestUseIndex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	indexes := []string{"index1", "index2"}
	qs := d.UseIndex(indexes...)
	if qs != d {
		t.Errorf("Expected %v, got %v", d, qs)
	}
}
