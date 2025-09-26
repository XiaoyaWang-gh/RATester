package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
	"github.com/gohugoio/hugo/hugofs"
)

func TestcollectDir_1(t *testing.T) {
	type args struct {
		dirPath  *paths.Path
		isDir    bool
		inFilter func(fim hugofs.FileMetaInfo) bool
	}
	tests := []struct {
		name    string
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

			c := &pagesCollector{}
			if err := c.collectDir(tt.args.dirPath, tt.args.isDir, tt.args.inFilter); (err != nil) != tt.wantErr {
				t.Errorf("pagesCollector.collectDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
