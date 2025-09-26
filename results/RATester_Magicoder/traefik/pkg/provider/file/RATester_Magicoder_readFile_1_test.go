package file

import (
	"fmt"
	"testing"
)

func TestreadFile_1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
		wantErr  bool
	}{
		{
			name:     "Test Case 1: Valid File",
			filename: "testdata/valid.txt",
			want:     "This is a valid file.",
			wantErr:  false,
		},
		{
			name:     "Test Case 2: Invalid File",
			filename: "testdata/invalid.txt",
			want:     "",
			wantErr:  true,
		},
		{
			name:     "Test Case 3: Empty File",
			filename: "",
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

			got, err := readFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
