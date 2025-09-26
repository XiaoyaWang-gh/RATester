package echo

import (
	"fmt"
	"testing"
)

func TestMustUint64_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var sourceParam string = "sourceParam"
	var dest *uint64 = nil
	var b *ValueBinder = &ValueBinder{}
	b.MustUint64(sourceParam, dest)
	if len(b.errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(b.errors))
	}
}
