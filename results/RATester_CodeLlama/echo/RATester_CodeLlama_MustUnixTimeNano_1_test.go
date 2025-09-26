package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestMustUnixTimeNano_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A test case
	sourceParam := "sourceParam"
	dest := &time.Time{}
	// and: The method to test
	b := &ValueBinder{}
	b.MustUnixTimeNano(sourceParam, dest)
	// when: Invoking the method
	// then: Expect the method to return without panicking
}
