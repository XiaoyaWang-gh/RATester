package logs

import (
	"fmt"
	"testing"
)

func TestReset_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	// when
	// then
}
