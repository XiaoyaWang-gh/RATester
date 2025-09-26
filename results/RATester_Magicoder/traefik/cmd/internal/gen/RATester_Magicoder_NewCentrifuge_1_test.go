package main

import (
	"fmt"
	"testing"
)

func TestNewCentrifuge_1(t *testing.T) {
	tests := []struct {
		name    string
		rootPkg string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			rootPkg: "github.com/traefik/traefik/v3/cmd/internal/gen",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			rootPkg: "invalid/package",
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

			_, err := NewCentrifuge(tt.rootPkg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCentrifuge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
