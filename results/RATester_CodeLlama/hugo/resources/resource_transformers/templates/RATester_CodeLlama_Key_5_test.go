package templates

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/internal"
)

func TestKey_5(t *testing.T) {
	tests := []struct {
		name string
		t    *executeAsTemplateTransform
		want internal.ResourceTransformationKey
	}{
		{
			name: "test",
			t: &executeAsTemplateTransform{
				targetPath: "test",
			},
			want: internal.NewResourceTransformationKey("execute-as-template", "test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.t.Key(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeAsTemplateTransform.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
