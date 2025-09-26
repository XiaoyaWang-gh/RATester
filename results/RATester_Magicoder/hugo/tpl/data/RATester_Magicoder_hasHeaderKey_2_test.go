package data

import (
	"fmt"
	"net/http"
	"testing"
)

func TesthasHeaderKey_2(t *testing.T) {
	header := make(http.Header)
	header.Add("Content-Type", "application/json")
	header.Add("Authorization", "Bearer token")

	tests := []struct {
		name string
		key  string
		want bool
	}{
		{
			name: "Key exists",
			key:  "Content-Type",
			want: true,
		},
		{
			name: "Key does not exist",
			key:  "Accept",
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

			if got := hasHeaderKey(header, tt.key); got != tt.want {
				t.Errorf("hasHeaderKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
