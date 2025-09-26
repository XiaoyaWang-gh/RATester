package acme

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-acme/lego/v4/certcrypto"
)

func TestNewAccount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	email := "test@example.com"
	keyTypeValue := "RSA"

	account, err := NewAccount(ctx, email, keyTypeValue)

	if err != nil {
		t.Errorf("NewAccount returned an error: %v", err)
	}

	if account.Email != email {
		t.Errorf("Expected email %s, but got %s", email, account.Email)
	}

	if account.KeyType != certcrypto.KeyType(keyTypeValue) {
		t.Errorf("Expected key type %s, but got %s", keyTypeValue, account.KeyType)
	}

	if account.PrivateKey == nil {
		t.Error("Private key is nil")
	}
}
