package echo

import (
	"fmt"
	"testing"
)

func TestMustUint32_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var sourceParam string = "sourceParam"
	var dest *uint32 = new(uint32)
	var b *ValueBinder = &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			return "123"
		},
	}
	b.MustUint32(sourceParam, dest)
	if *dest != 123 {
		t.Errorf("Expected *dest to be 123, but was %v", *dest)
	}
}
