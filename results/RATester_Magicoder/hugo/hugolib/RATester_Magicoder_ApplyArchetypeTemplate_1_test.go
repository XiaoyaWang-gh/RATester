package hugolib

import (
	"fmt"
	"io"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestApplyArchetypeTemplate_1(t *testing.T) {
	type args struct {
		w              io.Writer
		p              page.Page
		archetypeKind  string
		templateSource string
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

			if err := tt.f.ApplyArchetypeTemplate(tt.args.w, tt.args.p, tt.args.archetypeKind, tt.args.templateSource); (err != nil) != tt.wantErr {
				t.Errorf("ContentFactory.ApplyArchetypeTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
