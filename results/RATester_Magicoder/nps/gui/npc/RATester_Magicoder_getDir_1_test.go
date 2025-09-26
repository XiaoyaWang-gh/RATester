package main

import (
	"fmt"
	"testing"
)

func TestgetDir_1(t *testing.T) {
	tests := []struct {
		name    string
		wantDir string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			wantDir: "/data/data/org.nps.client/files",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			wantDir: "/home/user/.config",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotDir, err := getDir()
			if (err != nil) != tt.wantErr {
				t.Errorf("getDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDir != tt.wantDir {
				t.Errorf("getDir() = %v, want %v", gotDir, tt.wantDir)
			}
		})
	}
}
