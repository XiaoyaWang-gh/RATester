package types

import (
	"fmt"
	"testing"
)

func TestContains_5(t *testing.T) {
	q := &EvictingStringQueue{
		size: 5,
		vals: []string{"a", "b", "c", "d", "e"},
		set:  map[string]bool{"a": true, "b": true, "c": true, "d": true, "e": true},
	}

	tests := []struct {
		name string
		q    *EvictingStringQueue
		v    string
		want bool
	}{
		{
			name: "contains",
			q:    q,
			v:    "a",
			want: true,
		},
		{
			name: "does not contain",
			q:    q,
			v:    "f",
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

			if got := tt.q.Contains(tt.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
