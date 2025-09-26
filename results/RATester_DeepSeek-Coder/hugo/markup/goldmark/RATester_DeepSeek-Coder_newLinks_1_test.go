package goldmark

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/goldmark/goldmark_config"
	"github.com/yuin/goldmark"
)

func TestNewLinks_1(t *testing.T) {
	type args struct {
		cfg goldmark_config.Config
	}
	tests := []struct {
		name string
		args args
		want goldmark.Extender
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

			if got := newLinks(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
