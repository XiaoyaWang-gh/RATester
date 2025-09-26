package helpers

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
	"github.com/spf13/afero"
)

func TestWalk_2(t *testing.T) {
	type args struct {
		fs     afero.Fs
		root   string
		walker hugofs.WalkFunc
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

			if err := Walk(tt.args.fs, tt.args.root, tt.args.walker); (err != nil) != tt.wantErr {
				t.Errorf("Walk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
