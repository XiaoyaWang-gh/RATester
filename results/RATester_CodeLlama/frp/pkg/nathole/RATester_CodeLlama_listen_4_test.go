package nathole

import (
	"fmt"
	"testing"
)

func TestListen_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: add test cases
}
