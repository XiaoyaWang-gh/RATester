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
	origData := []byte("1234567890123456")
	crypted, err := AesEncrypt(origData, key)
	if err != nil {
		t.Errorf("AesEncrypt err: %v", err)
	}
	t.Logf("crypted: %v", crypted)
}
