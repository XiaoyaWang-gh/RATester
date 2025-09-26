package types

import (
	"fmt"
	"testing"
)

func TestPeek_4(t *testing.T) {
	q := &EvictingStringQueue{
		size: 5,
		vals: []string{"one", "two", "three", "four", "five"},
		set:  map[string]bool{"one": true, "two": true, "three": true, "four": true, "five": true},
	}

	tests := []struct {
		name string
		q    *EvictingStringQueue
		want string
	}{
		{
			name: "Peek from a non-empty queue",
			q:    q,
			want: "five",
		},
		{
			name: "Peek from an empty queue",
			q:    &EvictingStringQueue{},
			want: "",
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
