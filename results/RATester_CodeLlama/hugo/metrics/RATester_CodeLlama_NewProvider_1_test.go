package metrics

import (
	"fmt"
	"testing"
)

func TestNewProvider_1(t *testing.T) {
	t.Run("NewProvider", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Run("calculateHints=true", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := NewProvider(true)
			if p == nil {
				t.Error("NewProvider() returned nil")
			}
		})
		t.Run("calculateHints=false", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := NewProvider(false)
			if p == nil {
				t.Error("NewProvider() returned nil")
			}
		})
	})
}
