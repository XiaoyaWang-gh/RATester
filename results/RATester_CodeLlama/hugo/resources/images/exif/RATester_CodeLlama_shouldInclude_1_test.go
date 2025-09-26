package exif

import (
	"fmt"
	"testing"
)

func TestShouldInclude_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Decoder{}
	s := "test"
	if d.shouldInclude(s) {
		t.Errorf("shouldInclude() = %v, want %v", d.shouldInclude(s), false)
	}
}
