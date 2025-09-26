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
	crypted, _ := AesEncrypt([]byte("hello world"), key)
	origData, err := AesDecrypt(crypted, key)
	if err != nil {
		t.Errorf("AesDecrypt failed, err: %v", err)
		return
	}
	if string(origData) != "hello world" {
		t.Errorf("AesDecrypt failed, want: %s, got: %s", "hello world", string(origData))
		return
	}
	t.Logf("AesDecrypt success, want: %s, got: %s", "hello world", string(origData))
}
