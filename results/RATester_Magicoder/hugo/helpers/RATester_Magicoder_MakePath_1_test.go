package helpers

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/hugolib/paths"
)

func TestMakePath_1(t *testing.T) {
	type fields struct {
		Paths           *paths.Paths
		BaseFs          *filesystems.BaseFs
		ProcessingStats *ProcessingStats
		Fs              *hugofs.Fs
		Cfg             config.AllProvider
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
				Paths:           tt.fields.Paths,
				BaseFs:          tt.fields.BaseFs,
				ProcessingStats: tt.fields.ProcessingStats,
				Fs:              tt.fields.Fs,
				Cfg:             tt.fields.Cfg,
			}
			if got := p.MakePath(tt.args.s); got != tt.want {
				t.Errorf("MakePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
