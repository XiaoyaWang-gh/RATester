package crypt

import (
	"fmt"
	"testing"
)

func TestAesDecrypt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := []byte("1234567890123456")
	crypted := []byte("1234567890123456")
	origData, err := AesDecrypt(crypted, key)
	if err != nil {
		t.Errorf("AesDecrypt error: %v", err)
	}
	if string(origData) != "1234567890123456" {
		t.Errorf("AesDecrypt error: %v", origData)
	}
}
