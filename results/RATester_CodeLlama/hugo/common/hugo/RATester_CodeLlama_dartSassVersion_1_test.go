package hugo

import (
	"fmt"
	"testing"

	"github.com/bep/godartsass/v2"
)

func TestDartSassVersion_1(t *testing.T) {
	t.Parallel()
	t.Run("", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		if dartSassVersion() != (godartsass.DartSassVersion{}) {
			t.Errorf("got %v, want %v", dartSassVersion(), godartsass.DartSassVersion{})
		}
	})
	t.Run("", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		if !IsDartSassV2() {
			t.Errorf("got %v, want %v", IsDartSassV2(), true)
		}
	})
	t.Run("", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		if dartSassVersion() != (godartsass.DartSassVersion{}) {
			t.Errorf("got %v, want %v", dartSassVersion(), godartsass.DartSassVersion{})
		}
	})
}
