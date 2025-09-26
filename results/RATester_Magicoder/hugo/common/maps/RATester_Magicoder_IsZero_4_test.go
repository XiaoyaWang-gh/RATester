package maps

import (
	"fmt"
	"testing"
)

func TestIsZero_4(t *testing.T) {
	tests := []struct {
		name string
		p    Params
		want bool
	}{
		{
			name: "empty map",
			p:    Params{},
			want: true,
		},
		{
			name: "map with one key",
			p:    Params{MergeStrategyKey: "override"},
			want: false,
		},
		{
			name: "map with multiple keys",
			p:    Params{"key1": "value1", "key2": "value2"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.IsZero(); got != tt.want {
				t.Errorf("Params.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
