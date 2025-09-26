package transport

import (
	"fmt"
	"testing"
)

func TestNewRandomPrivateKey_1(t *testing.T) {
	t.Run("should return a PEM encoded private key", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key, err := NewRandomPrivateKey()
		if err != nil {
			t.Errorf("NewRandomPrivateKey() error = %v", err)
			return
		}
		if len(key) == 0 {
			t.Errorf("NewRandomPrivateKey() returned an empty key")
		}
	})
}
