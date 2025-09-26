package commands

import (
	"fmt"
	"testing"
)

func TestConvertJekyllPost_1(t *testing.T) {
	type args struct {
		path      string
		relPath   string
		targetDir string
		draft     bool
	}
	tests := []struct {
		name    string
		c       *importCommand
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.c.convertJekyllPost(tt.args.path, tt.args.relPath, tt.args.targetDir, tt.args.draft); (err != nil) != tt.wantErr {
				t.Errorf("convertJekyllPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
