package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
	"github.com/gohugoio/hugo/resources/page"
)

func TestwriteDestAlias_1(t *testing.T) {
	type args struct {
		path         string
		permalink    string
		outputFormat output.Format
		p            page.Page
	}
	tests := []struct {
		name    string
		s       *Site
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

			if err := tt.s.writeDestAlias(tt.args.path, tt.args.permalink, tt.args.outputFormat, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Site.writeDestAlias() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
