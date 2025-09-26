package generate

import (
	"crypto/rsa"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestPemCert_1(t *testing.T) {
	type args struct {
		privKey    *rsa.PrivateKey
		domain     string
		expiration time.Time
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

			got, err := PemCert(tt.args.privKey, tt.args.domain, tt.args.expiration)
			if (err != nil) != tt.wantErr {
				t.Errorf("PemCert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PemCert() = %v, want %v", got, tt.want)
			}
		})
	}
}
