package encryptcookie

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestGenerateKey_1(t *testing.T) {
	t.Run("should generate a key of the correct length", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		lengths := []int{16, 24, 32}

		for _, length := range lengths {
			key := GenerateKey(length)

			if len(key) != length {
				t.Errorf("expected key length of %d, got %d", length, len(key))
			}
		}
	})

	t.Run("should panic if length is not 16, 24, or 32", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic, but no panic occurred")
			}
		}()

		GenerateKey(40)
	})

	t.Run("should return a base64 encoded string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := GenerateKey(16)

		_, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			t.Errorf("expected key to be base64 encoded, but decoding failed: %v", err)
		}
	})
}
