package service

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClone_2(t *testing.T) {
	tests := []struct {
		name  string
		http  *http.Transport
		http2 *http.Transport
	}{
		{
			name:  "Test Clone",
			http:  &http.Transport{},
			http2: &http.Transport{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &smartRoundTripper{
				http:  tt.http,
				http2: tt.http2,
			}
			got := m.Clone()
			if got == nil {
				t.Errorf("Clone() = %v, want not nil", got)
			}
		})
	}
}
