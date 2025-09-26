package create

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestNewContent_1(t *testing.T) {
	type args struct {
		h          *hugolib.HugoSites
		kind       string
		targetPath string
		force      bool
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

			if err := NewContent(tt.args.h, tt.args.kind, tt.args.targetPath, tt.args.force); (err != nil) != tt.wantErr {
				t.Errorf("NewContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
