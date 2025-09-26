package csrf

import (
	"fmt"
	"testing"
)

func TestCompareTokens_1(t *testing.T) {
	t.Run("Testing equal tokens", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		token1 := []byte("token1")
		token2 := []byte("token1")
		result := compareTokens(token1, token2)
		if !result {
			t.Errorf("Expected true, got %v", result)
		}
	})

	t.Run("Testing unequal tokens", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		token1 := []byte("token1")
		token2 := []byte("token2")
		result := compareTokens(token1, token2)
		if result {
			t.Errorf("Expected false, got %v", result)
		}
	})
}
