package pagemeta

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestMarkupToMediaType_1(t *testing.T) {
	type args struct {
		s          string
		mediaTypes media.Types
	}
	tests := []struct {
		name string
		args args
		want media.Type
	}{
		{
			name: "test case 1",
			args: args{
				s:          "test case 1",
				mediaTypes: media.Types{},
			},
			want: media.Type{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MarkupToMediaType(tt.args.s, tt.args.mediaTypes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarkupToMediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}
