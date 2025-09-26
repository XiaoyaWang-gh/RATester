package modules

import (
	"fmt"
	"testing"
)

func TestListGoMods_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	// TODO(bep)
}
