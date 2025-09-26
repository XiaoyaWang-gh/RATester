package images

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yuin/goldmark"
)

func TestNew_62(t *testing.T) {
	tests := []struct {
		name string
		args bool
		want goldmark.Extender
	}{
		{
			name: "test case 1",
			args: true,
			want: &imagesExtension{wrapStandAloneImageWithinParagraph: true},
		},
		{
			name: "test case 2",
			args: false,
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

			if got := New(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
