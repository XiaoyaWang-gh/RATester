package template

import (
	"fmt"
	"testing"
)

func Test__1(t *testing.T) {
	t.Run("urlPartNone", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if urlPartNone != 0 {
			t.Errorf("Expected urlPartNone to be 0, got %d", urlPartNone)
		}
	})
}
