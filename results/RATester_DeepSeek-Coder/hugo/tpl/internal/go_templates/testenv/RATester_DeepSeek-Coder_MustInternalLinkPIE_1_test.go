package testenv

import (
	"fmt"
	"testing"
)

func TestMustInternalLinkPIE_1(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		MustInternalLinkPIE(t)
		if t.Failed() {
			t.Errorf("Expected MustInternalLinkPIE to pass")
		}
	})
}
