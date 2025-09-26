package mock

import (
	"fmt"
	"testing"
)

func TestUseIndex_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
