package web

import (
	"fmt"
	"testing"
)

func TestBuildTemplate_1(t *testing.T) {
	testCases := []struct {
		name    string
		dir     string
		files   []string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			dir:     "test_dir",
			files:   []string{"file1.tpl", "file2.tpl"},
			wantErr: false,
		},
		{
			name:    "Test case 2",
			dir:     "non_existing_dir",
			files:   []string{"file1.tpl", "file2.tpl"},
			wantErr: false,
		},
		{
			name:    "Test case 3",
			dir:     "test_dir",
			files:   []string{"non_existing_file.tpl"},
			wantErr: false,
		},
		{
			name:    "Test case 4",
			dir:     "test_dir",
			files:   []string{},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := BuildTemplate(tc.dir, tc.files...)
			if (err != nil) != tc.wantErr {
				t.Errorf("BuildTemplate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
