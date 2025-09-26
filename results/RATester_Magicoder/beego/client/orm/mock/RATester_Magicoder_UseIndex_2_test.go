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
	result := d.UseIndex(indexes...)
	if result != d {
		t.Errorf("Expected %v, but got %v", d, result)
	}
}
