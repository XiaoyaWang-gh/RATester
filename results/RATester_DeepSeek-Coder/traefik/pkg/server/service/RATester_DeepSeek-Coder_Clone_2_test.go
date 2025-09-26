package service

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClone_2(t *testing.T) {
	tests := []struct {
		name string
		sr   *smartRoundTripper
	}{
		{
			name: "Clone with nil http and http2",
			sr:   &smartRoundTripper{http: nil, http2: nil},
		},
		{
			name: "Clone with non-nil http and http2",
			sr:   &smartRoundTripper{http: &http.Transport{}, http2: &http.Transport{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.sr.Clone()
			if got == tt.sr {
				t.Errorf("Clone() = %v, want %v to be different", got, tt.sr)
			}
			if got.(*smartRoundTripper).http == tt.sr.http {
				t.Errorf("Clone() = %v, want %v to be different", got, tt.sr)
			}
			if got.(*smartRoundTripper).http2 == tt.sr.http2 {
				t.Errorf("Clone() = %v, want %v to be different", got, tt.sr)
			}
		})
	}
}
