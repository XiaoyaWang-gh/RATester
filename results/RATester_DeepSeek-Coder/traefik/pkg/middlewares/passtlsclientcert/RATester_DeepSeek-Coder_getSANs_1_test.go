package passtlsclientcert

import (
	"crypto/x509"
	"fmt"
	"reflect"
	"testing"
)

func TestGetSANs_1(t *testing.T) {
	type args struct {
		cert *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want []string
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

			if got := getSANs(tt.args.cert); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSANs() = %v, want %v", got, tt.want)
			}
		})
	}
}
