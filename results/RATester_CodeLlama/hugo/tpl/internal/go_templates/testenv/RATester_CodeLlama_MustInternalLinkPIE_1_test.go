package testenv

import (
	"fmt"
	"testing"
)

func TestMustInternalLinkPIE_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Helper()
	t.Parallel()
	MustInternalLinkPIE(t)
}
