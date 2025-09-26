package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefault_2(t *testing.T) {
	tests := []struct {
		name string
		opts []OptionFunc
		want *Engine
	}{
		{
			name: "Test with no options",
			opts: []OptionFunc{},
			want: &Engine{
				RedirectTrailingSlash:  true,
				RedirectFixedPath:      true,
				HandleMethodNotAllowed: true,
			},
		},
		{
			name: "Test with custom options",
			opts: []OptionFunc{
				func(e *Engine) {
					e.RedirectTrailingSlash = false
					e.RedirectFixedPath = false
					e.HandleMethodNotAllowed = false
				},
			},
			want: &Engine{
				RedirectTrailingSlash:  false,
				RedirectFixedPath:      false,
				HandleMethodNotAllowed: false,
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

			got := Default(tt.opts...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}
