package tables

import (
	"fmt"
	"testing"
)

func TestRenderHeaderOrRow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
