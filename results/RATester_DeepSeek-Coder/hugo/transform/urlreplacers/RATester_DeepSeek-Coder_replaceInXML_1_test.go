package urlreplacers

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/transform"
)

func TestReplaceInXML_1(t *testing.T) {
	type args struct {
		path string
		ct   transform.FromTo
	}
	tests := []struct {
		name string
		au   *absURLReplacer
		args args
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

			au := &absURLReplacer{
				htmlQuotes: tt.au.htmlQuotes,
				xmlQuotes:  tt.au.xmlQuotes,
			}
			au.replaceInXML(tt.args.path, tt.args.ct)
		})
	}
}
