package integrity

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/constants"
	"github.com/gohugoio/hugo/resources/internal"
)

func TestKey_11(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		algo string
		want internal.ResourceTransformationKey
	}{
		{
			name: "test1",
			algo: "test1",
			want: internal.NewResourceTransformationKey(constants.ResourceTransformationFingerprint, "test1"),
		},
		{
			name: "test2",
			algo: "test2",
			want: internal.NewResourceTransformationKey(constants.ResourceTransformationFingerprint, "test2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &fingerprintTransformation{
				algo: tt.algo,
			}
			if got := f.Key(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
