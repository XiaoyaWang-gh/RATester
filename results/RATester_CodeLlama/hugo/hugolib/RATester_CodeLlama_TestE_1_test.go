package hugolib

import (
	"fmt"
	"testing"
)

func TestTestE_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	if testing.Short() {
		t.Skip()
	}

	// TODO(bep)
}
