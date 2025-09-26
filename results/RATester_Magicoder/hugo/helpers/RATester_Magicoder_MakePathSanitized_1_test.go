package helpers

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/hugolib/paths"
)

func TestMakePathSanitized_1(t *testing.T) {
	type fields struct {
		Paths  *paths.Paths
		BaseFs *filesystems.BaseFs
		Cfg    config.AllProvider
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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

			p := &PathSpec{
				Paths:  tt.fields.Paths,
				BaseFs: tt.fields.BaseFs,
				Cfg:    tt.fields.Cfg,
			}
			if got := p.MakePathSanitized(tt.args.s); got != tt.want {
				t.Errorf("PathSpec.MakePathSanitized() = %v, want %v", got, tt.want)
			}
		})
	}
}
