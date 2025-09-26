package cert

import (
	"crypto/rsa"
	"fmt"
	"reflect"
	"testing"
)

func TestEncodePrivateKeyPEM_1(t *testing.T) {
	type args struct {
		key *rsa.PrivateKey
	}
	tests := []struct {
		name string
		args args
		want []byte
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

			if got := EncodePrivateKeyPEM(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodePrivateKeyPEM() = %v, want %v", got, tt.want)
			}
		})
	}
}
