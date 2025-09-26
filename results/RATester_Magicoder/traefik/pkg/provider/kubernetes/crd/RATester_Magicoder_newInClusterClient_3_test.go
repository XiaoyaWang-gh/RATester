package crd

import (
	"fmt"
	"testing"
)

func TestNewInClusterClient_3(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		wantErr  bool
	}{
		{
			name:     "Test Case 1",
			endpoint: "http://localhost:8080",
			wantErr:  false,
		},
		{
			name:     "Test Case 2",
			endpoint: "",
			wantErr:  false,
		},
		{
			name:     "Test Case 3",
			endpoint: "http://invalid-url",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := newInClusterClient(tt.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("newInClusterClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
