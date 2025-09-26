package generate

import (
	"fmt"
	"testing"
	"time"
)

func TestKeyPair_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	domain := "example.com"
	expiration := time.Now().Add(time.Hour * 24)

	certPEM, keyPEM, err := KeyPair(domain, expiration)
	if err != nil {
		t.Errorf("KeyPair() error = %v", err)
		return
	}

	if len(certPEM) == 0 {
		t.Errorf("KeyPair() got empty certPEM")
	}

	if len(keyPEM) == 0 {
		t.Errorf("KeyPair() got empty keyPEM")
	}
}
