package tables

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter/hooks"
	"github.com/gohugoio/hugo/markup/goldmark/internal/render"
)

func TestPeekTable_1(t *testing.T) {
	type args struct {
		ctx *render.Context
	}
	tests := []struct {
		name string
		args args
		want *hooks.Table
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

			r := &htmlRenderer{}
			if got := r.peekTable(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("peekTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
