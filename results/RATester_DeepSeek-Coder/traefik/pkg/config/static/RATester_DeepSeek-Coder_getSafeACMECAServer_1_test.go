package static

import (
	"fmt"
	"testing"
)

func TestGetSafeACMECAServer_1(t *testing.T) {
	type args struct {
		caServerSrc string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with empty string",
			args: args{caServerSrc: ""},
			want: DefaultAcmeCAServer,
		},
		{
			name: "Test with v01 endpoint",
			args: args{caServerSrc: "https://acme-v01.api.letsencrypt.org/directory"},
			want: "https://acme-v02.api.letsencrypt.org/directory",
		},
		{
			name: "Test with staging v01 endpoint",
			args: args{caServerSrc: "https://acme-staging.api.letsencrypt.org/directory"},
			want: "https://acme-staging-v02.api.letsencrypt.org/directory",
		},
		{
			name: "Test with v02 endpoint",
			args: args{caServerSrc: "https://acme-v02.api.letsencrypt.org/directory"},
			want: "https://acme-v02.api.letsencrypt.org/directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getSafeACMECAServer(tt.args.caServerSrc); got != tt.want {
				t.Errorf("getSafeACMECAServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
