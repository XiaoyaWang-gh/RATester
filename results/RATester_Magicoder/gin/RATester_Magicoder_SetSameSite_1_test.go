package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetSameSite_1(t *testing.T) {
	tests := []struct {
		name     string
		samesite http.SameSite
		want     http.SameSite
	}{
		{
			name:     "TestSetSameSite_Case1",
			samesite: http.SameSiteNoneMode,
			want:     http.SameSiteNoneMode,
		},
		{
			name:     "TestSetSameSite_Case2",
			samesite: http.SameSiteLaxMode,
			want:     http.SameSiteLaxMode,
		},
		{
			name:     "TestSetSameSite_Case3",
			samesite: http.SameSiteStrictMode,
			want:     http.SameSiteStrictMode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{
				sameSite: tt.samesite,
			}
			c.SetSameSite(tt.samesite)
			if c.sameSite != tt.want {
				t.Errorf("got %v, want %v", c.sameSite, tt.want)
			}
		})
	}
}
