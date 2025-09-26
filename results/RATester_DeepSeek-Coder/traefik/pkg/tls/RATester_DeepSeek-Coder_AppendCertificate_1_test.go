package tls

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestAppendCertificate_1(t *testing.T) {
	type args struct {
		certs     map[string]map[string]*tls.Certificate
		storeName string
	}
	tests := []struct {
		name    string
		c       *Certificate
		args    args
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

			if err := tt.c.AppendCertificate(tt.args.certs, tt.args.storeName); (err != nil) != tt.wantErr {
				t.Errorf("Certificate.AppendCertificate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
