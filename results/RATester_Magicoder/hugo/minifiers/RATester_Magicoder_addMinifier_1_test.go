package minifiers

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/media"
	"github.com/tdewolff/minify/v2"
)

func TestaddMinifier_1(t *testing.T) {
	type args struct {
		m      *minify.M
		mt     media.Types
		suffix string
		min    minify.Minifier
	}
	tests := []struct {
		name string
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

			addMinifier(tt.args.m, tt.args.mt, tt.args.suffix, tt.args.min)
		})
	}
}
