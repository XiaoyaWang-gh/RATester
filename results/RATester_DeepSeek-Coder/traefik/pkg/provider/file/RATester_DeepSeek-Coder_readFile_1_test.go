package file

import (
	"fmt"
	"testing"
)

func TestReadFile_1(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		want     string
		wantErr  bool
	}{
		{
			name:     "Test Case 1",
			filename: "test.txt",
			want:     "This is a test file.",
			wantErr:  false,
		},
		{
			name:     "Test Case 2",
			filename: "",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := readFile(tc.filename)
			if (err != nil) != tc.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("readFile() = %v, want %v", got, tc.want)
			}
		})
	}
}
