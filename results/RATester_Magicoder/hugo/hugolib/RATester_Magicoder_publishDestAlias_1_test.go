package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
	"github.com/gohugoio/hugo/resources/page"
)

func TestpublishDestAlias_1(t *testing.T) {
	type args struct {
		allowRoot    bool
		path         string
		permalink    string
		outputFormat output.Format
		p            page.Page
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

			s := &Site{}
			if err := s.publishDestAlias(tt.args.allowRoot, tt.args.path, tt.args.permalink, tt.args.outputFormat, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Site.publishDestAlias() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
