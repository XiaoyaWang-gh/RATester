package crypt

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAesEncrypt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := []byte("ABCDEFGHIJKLMNOP")
	origData := []byte("Hello, world")
	expectedCryptedData := []byte{0x8d, 0x8d, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c, 0x7c}

	cryptedData, err := AesEncrypt(origData, key)
	if err != nil {
		t.Errorf("AesEncrypt() error = %v", err)
		return
	}

	if !bytes.Equal(cryptedData, expectedCryptedData) {
		t.Errorf("AesEncrypt() = %v, want %v", cryptedData, expectedCryptedData)
	}
}
