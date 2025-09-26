package tailscale

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"
)

func TestIsValidCert_1(t *testing.T) {
	type args struct {
		cert   tls.Certificate
		domain string
		now    time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				cert:   tls.Certificate{},
				domain: "",
				now:    time.Time{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isValidCert(tt.args.cert, tt.args.domain, tt.args.now); got != tt.want {
				t.Errorf("isValidCert() = %v, want %v", got, tt.want)
			}
		})
	}
}
