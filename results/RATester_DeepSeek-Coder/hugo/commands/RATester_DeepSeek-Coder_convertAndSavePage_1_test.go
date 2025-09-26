package commands

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib"
	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/gohugoio/hugo/resources/page"
)

func TestConvertAndSavePage_1(t *testing.T) {
	type args struct {
		p            page.Page
		site         *hugolib.Site
		targetFormat metadecoders.Format
	}
	tests := []struct {
		name    string
		c       *convertCommand
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

			if err := tt.c.convertAndSavePage(tt.args.p, tt.args.site, tt.args.targetFormat); (err != nil) != tt.wantErr {
				t.Errorf("convertAndSavePage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
