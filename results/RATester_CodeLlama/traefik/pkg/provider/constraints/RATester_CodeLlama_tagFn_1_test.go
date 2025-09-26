package constraints

import (
	"fmt"
	"testing"
)

func TestTagFn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "name"
	tagFn := tagFn(name)
	tags := []string{"name", "age"}
	if !tagFn(tags) {
		t.Errorf("tagFn(%v) = %v, want %v", tags, false, true)
	}
	tags = []string{"age", "name"}
	if !tagFn(tags) {
		t.Errorf("tagFn(%v) = %v, want %v", tags, false, true)
	}
	tags = []string{"age"}
	if tagFn(tags) {
		t.Errorf("tagFn(%v) = %v, want %v", tags, true, false)
	}
	tags = []string{"name"}
	if tagFn(tags) {
		t.Errorf("tagFn(%v) = %v, want %v", tags, true, false)
	}
}
