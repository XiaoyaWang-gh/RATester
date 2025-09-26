package crypt

import (
	"fmt"
	"testing"
)

func TestAesEncrypt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := []byte("1234567890123456")
	origData := []byte("Hello, world!")

	crypted, err := AesEncrypt(origData, key)
	if err != nil {
		t.Errorf("AesEncrypt failed, err: %v", err)
		return
	}

	decrypted, err := AesDecrypt(crypted, key)
	if err != nil {
		t.Errorf("AesDecrypt failed, err: %v", err)
		return
	}

	if string(decrypted) != string(origData) {
		t.Errorf("AesEncrypt and AesDecrypt failed, decrypted: %s, origData: %s", decrypted, origData)
		return
	}
}
