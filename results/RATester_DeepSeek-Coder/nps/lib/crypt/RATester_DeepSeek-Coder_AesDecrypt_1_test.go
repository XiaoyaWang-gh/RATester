package crypt

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAesDecrypt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := []byte("ABCDEFGHIJKLMNOP")
	crypted, _ := hex.DecodeString("78787878787878787878787878787878")
	expected, _ := hex.DecodeString("3132333435363738393031323334353637383930")

	origData, err := AesDecrypt(crypted, key)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !bytes.Equal(origData, expected) {
		t.Errorf("Expected %v, got %v", expected, origData)
	}
}
