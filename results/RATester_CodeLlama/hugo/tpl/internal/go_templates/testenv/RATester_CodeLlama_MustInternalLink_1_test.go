package testenv

import (
	"fmt"
	"testing"
)

func TestMustInternalLink_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	MustInternalLink(t, true)
}
