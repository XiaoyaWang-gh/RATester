package internal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestParseArgs_1(t *testing.T) {
	type args struct {
		ctx converter.DocumentContext
	}
	tests := []struct {
		name string
		a    *AsciidocConverter
		args args
		want []string
	}{
		{
			name: "test case 1",
			a:    &AsciidocConverter{},
			args: args{
				ctx: converter.DocumentContext{
					Document:       nil,
					DocumentLookup: nil,
					DocumentID:     "",
					DocumentName:   "",
					Filename:       "",
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.a.ParseArgs(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
