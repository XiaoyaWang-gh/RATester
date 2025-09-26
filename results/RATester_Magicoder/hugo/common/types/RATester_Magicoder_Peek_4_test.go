package types

import (
	"fmt"
	"testing"
)

func TestPeek_4(t *testing.T) {
	q := &EvictingStringQueue{
		size: 3,
		vals: []string{"a", "b", "c"},
		set:  map[string]bool{"a": true, "b": true, "c": true},
	}

	tests := []struct {
		name string
		q    *EvictingStringQueue
		want string
	}{
		{
			name: "Test Peek",
			q:    q,
			want: "c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.q.Peek(); got != tt.want {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
