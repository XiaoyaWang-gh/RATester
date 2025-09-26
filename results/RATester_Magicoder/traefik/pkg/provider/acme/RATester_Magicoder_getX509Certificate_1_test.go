package acme

import (
	"context"
	"crypto/x509"
	"fmt"
	"reflect"
	"testing"
)

func TestGetX509Certificate_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		cert *Certificate
	}
	tests := []struct {
		name    string
		args    args
		want    *x509.Certificate
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

			got, err := getX509Certificate(tt.args.ctx, tt.args.cert)
			if (err != nil) != tt.wantErr {
				t.Errorf("getX509Certificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getX509Certificate() = %v, want %v", got, tt.want)
			}
		})
	}
}
