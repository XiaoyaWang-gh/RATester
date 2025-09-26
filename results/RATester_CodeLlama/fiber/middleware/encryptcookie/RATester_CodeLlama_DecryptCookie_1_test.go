package encryptcookie

import (
	"fmt"
	"testing"
)

func TestDecryptCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	value := "encrypted value"
	key := "key"

	_, err := DecryptCookie(value, key)
	if err != nil {
		t.Errorf("failed to decrypt cookie: %v", err)
	}
}
