package echo

import (
	"fmt"
	"testing"
)

func TestBindWithDelimiter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	sourceParam := "sourceParam"
	dest := []string{}
	delimiter := "delimiter"
	valueMustExist := true
	b := &ValueBinder{}
	// when
	b.bindWithDelimiter(sourceParam, dest, delimiter, valueMustExist)
	// then
	// TODO: add assertions
}
