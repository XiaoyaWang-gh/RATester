package install

import (
	"fmt"
	"testing"
)

func TestChMod_1(t *testing.T) {
	t.Run("Test chMod", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// TODO: Add test cases.
	})
}
