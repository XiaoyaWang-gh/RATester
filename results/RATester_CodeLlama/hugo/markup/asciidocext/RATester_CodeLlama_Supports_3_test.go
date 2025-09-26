package asciidocext

import (
	"fmt"
	"testing"
)

func TestSupports_3(t *testing.T) {
	t.Run("Supports", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if !Supports() {
			t.Error("expected true")
		}
	})
}
