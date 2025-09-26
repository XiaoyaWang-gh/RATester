package commands

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestStat_3(t *testing.T) {
	testCases := []struct {
		name     string
		fileName string
		wantErr  bool
	}{
		{
			name:     "File exists",
			fileName: "test.txt",
			wantErr:  false,
		},
		{
			name:     "File does not exist",
			fileName: "nonexistent.txt",
			wantErr:  true,
		},
	}

	fs := &countingStatFs{
		Fs:          afero.NewMemMapFs(),
		statCounter: 0,
	}

	// Create test files
	afero.WriteFile(fs, "test.txt", []byte("test content"), 0644)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := fs.Stat(tc.fileName)
			if (err != nil) != tc.wantErr {
				t.Errorf("Stat() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
