package cssjs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecodePostCSSOptions_1(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]any
		want PostCSSOptions
	}{
		{
			name: "empty",
			m:    map[string]any{},
			want: PostCSSOptions{},
		},
		{
			name: "no-map",
			m: map[string]any{
				"no-map": true,
			},
			want: PostCSSOptions{
				NoMap: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			opts, err := decodePostCSSOptions(tt.m)
			if (err != nil) != (tt.want.NoMap != tt.want.NoMap) {
				t.Errorf("decodePostCSSOptions() error = %v, wantErr %v", err, tt.want.NoMap != tt.want.NoMap)
				return
			}
			if !reflect.DeepEqual(opts, tt.want) {
				t.Errorf("decodePostCSSOptions() = %v, want %v", opts, tt.want)
			}
		})
	}
}
