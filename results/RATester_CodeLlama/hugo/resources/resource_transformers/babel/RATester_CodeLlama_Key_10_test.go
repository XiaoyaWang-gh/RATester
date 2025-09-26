package babel

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/internal"
)

func TestKey_10(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		t    *babelTransformation
		want internal.ResourceTransformationKey
	}{
		{
			name: "test",
			t: &babelTransformation{
				options: Options{
					Config: "test",
				},
			},
			want: internal.NewResourceTransformationKey("babel", Options{
				Config: "test",
			}),
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
				t.Errorf("babelTransformation.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
