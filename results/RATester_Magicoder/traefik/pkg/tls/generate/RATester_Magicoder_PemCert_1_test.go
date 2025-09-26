package generate

import (
	"crypto/rsa"
	"fmt"
	"math/big"
	"testing"
	"time"
)

func TestPemCert_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	privKey := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: big.NewInt(1),
			E: 1,
		},
		D:           big.NewInt(1),
		Primes:      []*big.Int{big.NewInt(1)},
		Precomputed: rsa.PrecomputedValues{},
	}
	domain := "example.com"
	expiration := time.Now().Add(time.Hour)

	_, err := PemCert(privKey, domain, expiration)
	if err != nil {
		t.Errorf("PemCert() error = %v", err)
	}
}
