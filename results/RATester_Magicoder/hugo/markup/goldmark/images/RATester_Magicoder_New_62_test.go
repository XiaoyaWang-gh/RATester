package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark"
)

func TestNew_62(t *testing.T) {
	type args struct {
		wrapStandAloneImageWithinParagraph bool
	}
	tests := []struct {
		name string
		args args
		want goldmark.Extender
	}{
		{
			name: "Test case 1",
			args: args{
				wrapStandAloneImageWithinParagraph: true,
			},
			want: &imagesExtension{wrapStandAloneImageWithinParagraph: true},
		},
		{
			name: "Test case 2",
			args: args{
				wrapStandAloneImageWithinParagraph: false,
			},
			want: &imagesExtension{wrapStandAloneImageWithinParagraph: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := New(tt.args.wrapStandAloneImageWithinParagraph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
