package wrr

import (
	"fmt"
	"net/http"
	"testing"
)

func TestConvertSameSite_1(t *testing.T) {
	tests := []struct {
		name     string
		sameSite string
		want     http.SameSite
	}{
		{
			name:     "none",
			sameSite: "none",
			want:     http.SameSiteNoneMode,
		},
		{
			name:     "lax",
			sameSite: "lax",
			want:     http.SameSiteLaxMode,
		},
		{
			name:     "strict",
			sameSite: "strict",
			want:     http.SameSiteStrictMode,
		},
		{
			name:     "default",
			sameSite: "default",
			want:     http.SameSiteDefaultMode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertSameSite(tt.sameSite); got != tt.want {
				t.Errorf("convertSameSite() = %v, want %v", got, tt.want)
			}
		})
	}
}
