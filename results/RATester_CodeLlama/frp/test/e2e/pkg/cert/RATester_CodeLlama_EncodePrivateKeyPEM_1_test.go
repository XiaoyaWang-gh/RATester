package cert

import (
	"crypto/rsa"
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestEncodePrivateKeyPEM_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	key := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: big.NewInt(123),
			E: 123,
		},
		D: big.NewInt(123),
	}
	want := []byte("-----BEGIN RSA PRIVATE KEY-----\n" +
		"MIIEpAIBAAKCAQEAy123\n" +
		"-----END RSA PRIVATE KEY-----\n")
	got := EncodePrivateKeyPEM(key)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("EncodePrivateKeyPEM() = %v, want %v", got, want)
	}
}
