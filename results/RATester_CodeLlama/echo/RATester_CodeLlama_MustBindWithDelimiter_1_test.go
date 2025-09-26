package echo

import (
	"fmt"
	"testing"
)

func TestMustBindWithDelimiter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	sourceParam := "sourceParam"
	dest := "dest"
	delimiter := "delimiter"
	b := &ValueBinder{}
	// when
	b.MustBindWithDelimiter(sourceParam, dest, delimiter)
	// then
	// TODO: add assertions
}
