package exif

import (
	"fmt"
	"testing"
)

func TestWithDateDisabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Decoder{}
	WithDateDisabled(true)(d)
	if !d.noDate {
		t.Error("WithDateDisabled failed to set noDate to true")
	}
}
