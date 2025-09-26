package utils

import (
	"fmt"
	"testing"
)

func TestSearchFile_1(t *testing.T) {
	type test struct {
		name     string
		filename string
		paths    []string
		want     string
		wantErr  bool
	}

	tests := []test{
		{
			name:     "File exists",
			filename: "test.txt",
			paths:    []string{"/home/user", "/usr/local"},
			want:     "/home/user/test.txt",
			wantErr:  false,
		},
		{
			name:     "File does not exist",
			filename: "nonexistent.txt",
			paths:    []string{"/home/user", "/usr/local"},
			want:     "",
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

			got, err := SearchFile(tt.filename, tt.paths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SearchFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
