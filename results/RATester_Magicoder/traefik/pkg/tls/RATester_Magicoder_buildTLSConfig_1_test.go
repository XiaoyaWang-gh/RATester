package tls

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"
)

func TestBuildTLSConfig_1(t *testing.T) {
	tests := []struct {
		name      string
		tlsOption Options
		want      *tls.Config
		wantErr   bool
	}{
		{
			name: "Test Case 1",
			tlsOption: Options{
				MinVersion: "TLS1.2",
			},
			want: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			tlsOption: Options{
				MinVersion: "TLS1.3",
			},
			want: &tls.Config{
				MinVersion: tls.VersionTLS13,
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			tlsOption: Options{
				MinVersion: "TLS1.4",
			},
			want:    nil,
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

			got, err := buildTLSConfig(tt.tlsOption)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildTLSConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTLSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
