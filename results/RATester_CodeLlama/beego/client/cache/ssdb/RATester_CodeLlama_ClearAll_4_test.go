package ssdb

import (
	"fmt"
	"testing"
)

func TestClearAll_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
