package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/resources/page"
)

func TestsetMetaPost_1(t *testing.T) {
	type args struct {
		cascade map[page.PageMatcher]maps.Params
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

			ps := &pageState{}
			if err := ps.setMetaPost(tt.args.cascade); (err != nil) != tt.wantErr {
				t.Errorf("setMetaPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
