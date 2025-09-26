package mock

import (
	"fmt"
	"testing"
)

func TestIgnoreIndex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DoNothingQuerySetter{}
	indexes := []string{"idx1", "idx2"}
	result := d.IgnoreIndex(indexes...)
	if result != d {
		t.Errorf("Expected IgnoreIndex to return the original QuerySetter, but got %v", result)
	}
}
