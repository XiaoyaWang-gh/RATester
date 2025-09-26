package commands

import (
	"fmt"
	"testing"
)

func TestconvertJekyllPost_1(t *testing.T) {
	type args struct {
		path      string
		relPath   string
		targetDir string
		draft     bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				path:      "testdata/test.md",
				relPath:   "test.md",
				targetDir: "testdata/output",
				draft:     false,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				path:      "testdata/non-existent.md",
				relPath:   "non-existent.md",
				targetDir: "testdata/output",
				draft:     false,
			},
			wantErr: true,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &importCommand{}
			if err := c.convertJekyllPost(tt.args.path, tt.args.relPath, tt.args.targetDir, tt.args.draft); (err != nil) != tt.wantErr {
				t.Errorf("convertJekyllPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
