package types

import (
	"fmt"
	"testing"
)

func TestContains_5(t *testing.T) {
	q := &EvictingStringQueue{
		size: 3,
		vals: []string{"a", "b", "c"},
		set:  map[string]bool{"a": true, "b": true, "c": true},
	}

	tests := []struct {
		name string
		q    *EvictingStringQueue
		arg  string
		want bool
	}{
		{
			name: "contains",
			q:    q,
			arg:  "a",
			want: true,
		},
		{
			name: "does not contain",
			q:    q,
			arg:  "d",
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

			if got := tt.q.Contains(tt.arg); got != tt.want {
				t.Errorf("EvictingStringQueue.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
