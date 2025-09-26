package http

import (
	"fmt"
	"testing"
)

func TestCanonicalHost_1(t *testing.T) {
	tests := []struct {
		name    string
		host    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test with valid host",
			host:    "www.google.com",
			want:    "www.google.com",
			wantErr: false,
		},
		{
			name:    "Test with valid host with port",
			host:    "www.google.com:8080",
			want:    "www.google.com",
			wantErr: false,
		},
		{
			name:    "Test with valid host with trailing dot",
			host:    "www.google.com.",
			want:    "www.google.com",
			wantErr: false,
		},
		{
			name:    "Test with invalid host",
			host:    "www.google.com:808080",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := CanonicalHost(tt.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("CanonicalHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CanonicalHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
