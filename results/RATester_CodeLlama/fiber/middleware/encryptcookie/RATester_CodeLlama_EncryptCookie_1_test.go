package encryptcookie

import (
	"fmt"
	"testing"
)

func TestEncryptCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	value := "test"
	key := "test"
	encryptedValue, err := EncryptCookie(value, key)
	if err != nil {
		t.Errorf("EncryptCookie() failed: %v", err)
	}
	if encryptedValue != "test" {
		t.Errorf("EncryptCookie() failed: %v", encryptedValue)
	}
}
