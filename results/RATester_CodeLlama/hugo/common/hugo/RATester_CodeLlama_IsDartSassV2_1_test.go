package hugo

import (
	"fmt"
	"testing"
)

func TestIsDartSassV2_1(t *testing.T) {
	t.Parallel()

	t.Run("true", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		if !IsDartSassV2() {
			t.Error("expected true")
		}
	})

	t.Run("false", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		if IsDartSassV2() {
			t.Error("expected false")
		}
	})
}
