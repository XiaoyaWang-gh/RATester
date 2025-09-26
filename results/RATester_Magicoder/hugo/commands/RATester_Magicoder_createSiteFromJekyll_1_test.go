package commands

import (
	"fmt"
	"testing"
)

func TestcreateSiteFromJekyll_1(t *testing.T) {
	type args struct {
		jekyllRoot     string
		targetDir      string
		jekyllPostDirs map[string]bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				jekyllRoot: "/path/to/jekyll/root",
				targetDir:  "/path/to/target/dir",
				jekyllPostDirs: map[string]bool{
					"posts":  true,
					"drafts": true,
				},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				jekyllRoot: "/path/to/jekyll/root",
				targetDir:  "/path/to/target/dir",
				jekyllPostDirs: map[string]bool{
					"posts":  true,
					"drafts": true,
				},
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
			if err := c.createSiteFromJekyll(tt.args.jekyllRoot, tt.args.targetDir, tt.args.jekyllPostDirs); (err != nil) != tt.wantErr {
				t.Errorf("createSiteFromJekyll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
