package minifier

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/internal"
)

func TestKey_6(t *testing.T) {
	tt := []struct {
		name string
		key  internal.ResourceTransformationKey
		want string
	}{
		{
			name: "TestKey",
			key:  internal.NewResourceTransformationKey("minify"),
			want: "minify",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.key.Value()
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
