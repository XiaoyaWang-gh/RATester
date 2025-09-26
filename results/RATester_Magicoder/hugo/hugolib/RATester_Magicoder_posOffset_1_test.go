package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestposOffset_1(t *testing.T) {
	type args struct {
		offset int
	}
	tests := []struct {
		name string
		p    *pageState
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

			if got := tt.p.posOffset(tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("posOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
