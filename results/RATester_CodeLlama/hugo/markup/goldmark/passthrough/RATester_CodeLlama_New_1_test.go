package passthrough

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/goldmark/goldmark_config"
	"github.com/yuin/goldmark"
)

func TestNew_1(t *testing.T) {
	tests := []struct {
		name string
		cfg  goldmark_config.Passthrough
		want goldmark.Extender
	}{
		{
			name: "nil",
			cfg:  goldmark_config.Passthrough{},
			want: nil,
		},
		{
			name: "enabled",
			cfg:  goldmark_config.Passthrough{Enable: true},
			want: &passthroughExtension{cfg: goldmark_config.Passthrough{Enable: true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := New(tt.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
