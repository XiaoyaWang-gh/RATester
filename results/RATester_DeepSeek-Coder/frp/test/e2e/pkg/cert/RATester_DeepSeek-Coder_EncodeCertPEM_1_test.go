package cert

import (
	"crypto/x509"
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeCertPEM_1(t *testing.T) {
	type args struct {
		ct *x509.Certificate
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

			if got := EncodeCertPEM(tt.args.ct); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeCertPEM() = %v, want %v", got, tt.want)
			}
		})
	}
}
