package hugolib

import (
	"fmt"
	"io"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/resources/page"
)

func TestApplyArchetypeFi_1(t *testing.T) {
	type args struct {
		w             io.Writer
		p             page.Page
		archetypeKind string
		fi            hugofs.FileMetaInfo
	}
	tests := []struct {
		name    string
		f       ContentFactory
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

			if err := tt.f.ApplyArchetypeFi(tt.args.w, tt.args.p, tt.args.archetypeKind, tt.args.fi); (err != nil) != tt.wantErr {
				t.Errorf("ContentFactory.ApplyArchetypeFi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
