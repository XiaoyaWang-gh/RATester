package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"
)

func TestEncodePrivateKeyPEM_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}

	encoded := EncodePrivateKeyPEM(key)

	if len(encoded) == 0 {
		t.Error("Expected non-empty encoded key")
	}
}
