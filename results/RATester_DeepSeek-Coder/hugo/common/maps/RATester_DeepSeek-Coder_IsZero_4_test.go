package maps

import (
	"fmt"
	"testing"
)

func TestIsZero_4(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    Params
		want bool
	}{
		{
			name: "empty",
			p:    Params{},
			want: true,
		},
		{
			name: "single key",
			p:    Params{MergeStrategyKey: "override"},
			want: false,
		},
		{
			name: "multiple keys",
			p:    Params{"a": "1", "b": "2", MergeStrategyKey: "override"},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt // NOTE: https://golang.org/doc/faq#closures_and_goroutines
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			if got := tt.p.IsZero(); got != tt.want {
				t.Errorf("Params.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
