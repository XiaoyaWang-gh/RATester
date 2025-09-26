package crd

import (
	"fmt"
	"testing"
)

func TestParseServiceProtocol_1(t *testing.T) {
	tests := []struct {
		name           string
		providedScheme string
		portName       string
		portNumber     int32
		want           string
		wantErr        bool
	}{
		{
			name:           "http",
			providedScheme: httpProtocol,
			portName:       "",
			portNumber:     0,
			want:           httpProtocol,
			wantErr:        false,
		},
		{
			name:           "https",
			providedScheme: "",
			portName:       "https",
			portNumber:     443,
			want:           httpsProtocol,
			wantErr:        false,
		},
		{
			name:           "h2c",
			providedScheme: "h2c",
			portName:       "",
			portNumber:     0,
			want:           "h2c",
			wantErr:        false,
		},
		{
			name:           "invalid",
			providedScheme: "invalid",
			portName:       "",
			portNumber:     0,
			want:           "",
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := parseServiceProtocol(tt.providedScheme, tt.portName, tt.portNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseServiceProtocol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseServiceProtocol() = %v, want %v", got, tt.want)
			}
		})
	}
}
