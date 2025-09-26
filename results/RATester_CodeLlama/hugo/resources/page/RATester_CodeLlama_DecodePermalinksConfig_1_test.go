package page

import (
	"fmt"
	"testing"
)

func TestDecodePermalinksConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// TODO: add test cases
}
