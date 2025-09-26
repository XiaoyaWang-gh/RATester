package exif

import (
	"fmt"
	"testing"
)

func TestWithLatLongDisabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	decoder := &Decoder{}
	WithLatLongDisabled(true)(decoder)
	if !decoder.noLatLong {
		t.Error("Expected WithLatLongDisabled to set noLatLong to true")
	}
}
