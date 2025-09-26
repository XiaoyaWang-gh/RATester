package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestnewPageOutput_1(t *testing.T) {
	type args struct {
		ps     *pageState
		pp     pagePaths
		f      output.Format
		render bool
	}
	tests := []struct {
		name string
		args args
		want *pageOutput
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

			if got := newPageOutput(tt.args.ps, tt.args.pp, tt.args.f, tt.args.render); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPageOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
