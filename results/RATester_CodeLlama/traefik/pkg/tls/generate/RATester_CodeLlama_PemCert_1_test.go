package generate

import (
	"crypto/rsa"
	"fmt"
	"testing"
	"time"
)

func TestPemCert_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	privKey := &rsa.PrivateKey{}
	// CONTEXT_1
	domain := "example.com"
	// CONTEXT_2
	expiration := time.Now()
	// CONTEXT_3
	pemCert, err := PemCert(privKey, domain, expiration)
	if err != nil {
		t.Fatal(err)
	}
	// CONTEXT_4
	if len(pemCert) == 0 {
		t.Fatal("pemCert is empty")
	}
	// CONTEXT_5
	if expiration.IsZero() {
		t.Fatal("expiration is zero")
	}
	// CONTEXT_6
	if expiration.After(time.Now()) {
		t.Fatal("expiration is after now")
	}
	// CONTEXT_7
	if err != nil {
		t.Fatal(err)
	}
}
