package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestResolvePosition_1(t *testing.T) {
	type args struct {
		ctx any
	}
	tests := []struct {
		name string
		hr   hookRendererTemplate
		args args
		want text.Position
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

			if got := tt.hr.ResolvePosition(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolvePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
