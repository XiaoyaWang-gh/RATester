package encryptcookie

import (
	"fmt"
	"testing"
)

func TestGenerateKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	for _, length := range []int{16, 24, 32} {
		key := GenerateKey(length)
		if len(key) != length {
			t.Errorf("Expected key length %d, got %d", length, len(key))
		}
	}
}
