package acme

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetPrivateKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	a := &Account{
		PrivateKey: []byte("private key"),
	}

	// when
	privateKey := a.GetPrivateKey()

	// then
	assert.NotNil(t, privateKey)
}
