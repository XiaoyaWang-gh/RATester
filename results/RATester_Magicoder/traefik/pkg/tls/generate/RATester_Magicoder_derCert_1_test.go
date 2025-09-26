package generate

import (
	"crypto/rsa"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDerCert_1(t *testing.T) {
	type args struct {
		privKey    *rsa.PrivateKey
		expiration time.Time
		domain     string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
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

			got, err := derCert(tt.args.privKey, tt.args.expiration, tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("derCert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("derCert() = %v, want %v", got, tt.want)
			}
		})
	}
}
