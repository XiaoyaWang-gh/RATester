package passtlsclientcert

import (
	"context"
	"crypto/x509"
	"fmt"
	"testing"
)

func TestExtractCertificate_1(t *testing.T) {
	ctx := context.Background()

	cert := &x509.Certificate{
		Raw: []byte{0x00, 0x01, 0x02, 0x03},
	}

	tests := []struct {
		name string
		ctx  context.Context
		cert *x509.Certificate
		want string
	}{
		{
			name: "Success",
			ctx:  ctx,
			cert: cert,
			want: "-----BEGIN CERTIFICATE-----\nABCDEF==\n-----END CERTIFICATE-----\n",
		},
		{
			name: "Fail",
			ctx:  ctx,
			cert: nil,
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

			if got := extractCertificate(tt.ctx, tt.cert); got != tt.want {
				t.Errorf("extractCertificate() = %v, want %v", got, tt.want)
			}
		})
	}
}
