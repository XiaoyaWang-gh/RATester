package logs

import (
	"fmt"
	"testing"
)

func TestFlush_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO
}
