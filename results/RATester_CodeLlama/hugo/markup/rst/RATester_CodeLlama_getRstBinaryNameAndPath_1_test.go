package rst

import (
	"fmt"
	"testing"
)

func TestGetRstBinaryNameAndPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
