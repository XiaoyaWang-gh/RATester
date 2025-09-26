package streamserver

import (
	"fmt"
	"testing"
)

func TestWithCustomHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// TODO:
}
