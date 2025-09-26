package acme

import (
	"crypto"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/registration"
)

func TestGetPrivateKey_1(t *testing.T) {
	type fields struct {
		Email        string
		Registration *registration.Resource
		PrivateKey   []byte
		KeyType      certcrypto.KeyType
	}
	tests := []struct {
		name   string
		fields fields
		want   crypto.PrivateKey
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			a := &Account{
				Email:        tt.fields.Email,
				Registration: tt.fields.Registration,
				PrivateKey:   tt.fields.PrivateKey,
				KeyType:      tt.fields.KeyType,
			}
			if got := a.GetPrivateKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.GetPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
