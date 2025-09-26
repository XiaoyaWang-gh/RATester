package debug

import (
	"fmt"
	"sync"
	"testing"
)

func TestVisualizeSpaces_1(t *testing.T) {
	ns := &Namespace{
		timersMu: sync.Mutex{},
		timers:   make(map[string][]*timer),
	}

	type test struct {
		name string
		val  any
		want string
	}

	tests := []test{
		{
			name: "Test case 1",
			val:  "Hello, World!",
			want: "Hello, World!",
		},
		{
			name: "Test case 2",
			val:  "This is a test",
			want: "This is a test",
		},
		{
			name: "Test case 3",
			val:  "This is another test",
			want: "This is another test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ns.VisualizeSpaces(tt.val); got != tt.want {
				t.Errorf("VisualizeSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
