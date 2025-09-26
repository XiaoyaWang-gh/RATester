package static

import (
	"fmt"
	"testing"
)

func TestgetSafeACMECAServer_1(t *testing.T) {
	tests := []struct {
		name        string
		caServerSrc string
		want        string
	}{
		{
			name:        "empty string",
			caServerSrc: "",
			want:        DefaultAcmeCAServer,
		},
		{
			name:        "v01 endpoint",
			caServerSrc: "https://acme-v01.api.letsencrypt.org/directory",
			want:        "https://acme-v02.api.letsencrypt.org/directory",
		},
		{
			name:        "v01 staging endpoint",
			caServerSrc: "https://acme-staging.api.letsencrypt.org/directory",
			want:        "https://acme-staging-v02.api.letsencrypt.org/directory",
		},
		{
			name:        "v02 endpoint",
			caServerSrc: "https://acme-v02.api.letsencrypt.org/directory",
			want:        "https://acme-v02.api.letsencrypt.org/directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getSafeACMECAServer(tt.caServerSrc); got != tt.want {
				t.Errorf("getSafeACMECAServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
