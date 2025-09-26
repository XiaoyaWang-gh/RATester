package commands

import (
	"fmt"
	"testing"
)

func TestCopyJekyllFilesAndFolders_1(t *testing.T) {
	type args struct {
		jekyllRoot     string
		dest           string
		jekyllPostDirs map[string]bool
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

			c := &importCommand{}
			if err := c.copyJekyllFilesAndFolders(tt.args.jekyllRoot, tt.args.dest, tt.args.jekyllPostDirs); (err != nil) != tt.wantErr {
				t.Errorf("importCommand.copyJekyllFilesAndFolders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
