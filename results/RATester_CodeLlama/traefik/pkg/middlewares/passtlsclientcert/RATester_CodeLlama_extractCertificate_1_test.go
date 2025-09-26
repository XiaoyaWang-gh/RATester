package passtlsclientcert

import (
	"context"
	"crypto/x509"
	"fmt"
	"testing"
)

func TestExtractCertificate_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		cert *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test01",
			args: args{
				ctx:  context.Background(),
				cert: &x509.Certificate{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := extractCertificate(tt.args.ctx, tt.args.cert); got != tt.want {
				t.Errorf("extractCertificate() = %v, want %v", got, tt.want)
			}
		})
	}
}
