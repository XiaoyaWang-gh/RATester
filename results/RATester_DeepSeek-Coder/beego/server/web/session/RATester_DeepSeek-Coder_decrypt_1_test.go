package session

import (
	"crypto/aes"
	"fmt"
	"testing"
)

func TestDecrypt_1(t *testing.T) {
	t.Run("Test decrypt function", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		key := []byte("thisisaverysecurekey")
		block, err := aes.NewCipher(key)
		if err != nil {
			t.Errorf("Error creating new cipher: %v", err)
		}

		value := []byte("thisisatestvalue")
		encryptedValue, err := encrypt(block, value)
		if err != nil {
			t.Errorf("Error encrypting value: %v", err)
		}

		decryptedValue, err := decrypt(block, encryptedValue)
		if err != nil {
			t.Errorf("Error decrypting value: %v", err)
		}

		if string(decryptedValue) != string(value) {
			t.Errorf("Expected %s, got %s", value, decryptedValue)
		}
	})
}
